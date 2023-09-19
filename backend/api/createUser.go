package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
	"github.com/kwandapchumba/bookmarkmonster/vultr"
)

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *createUserRequest

	if err := jsonDecoder.Decode(&req); err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		} else {
			log.Printf("error decoding request body to struct: %v", err)
			utils.Response(w, "something went wrong", http.StatusBadRequest)
			return
		}
	}

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		log.Printf("could not create new pool: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	defer pool.Close()

	q := sqlc.New(pool)

	link, err := q.GetEmailConfirmationLink(ctx, req.Email)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Println("email confirmation link does not exist")
			utils.Response(w, "email confirmation link does not exist", http.StatusNotFound)
			return
		default:
			log.Printf("could not get email confirmation link: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}
	}

	if time.Now().UTC().After(link.Expiry) {
		log.Println("email confirmation token has expired")
		utils.Response(w, "email confirmation token has expired", http.StatusForbidden)
		return
	}

	if link.Code != req.Token {
		log.Printf("tokens %s and %s do not match", req.Token, link.Code)
		utils.Response(w, "invalid email verification token", http.StatusForbidden)
		return
	}

	bg := "213555"
	color := "FFFFFF"

	url := fmt.Sprintf("https://ui-avatars.com/api/?name=%s&background=%s&color=%s&bold=true&length=1", req.Email, bg, color)

	response, err := http.Get(url)
	if err != nil {
		log.Printf("could not get user profile picture from name: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	defer response.Body.Close()

	profilePicFile, err := os.Create("profilePic.png")
	if err != nil {
		log.Printf("could not create favicon file: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	defer profilePicFile.Close()

	_, err = io.Copy(profilePicFile, response.Body)
	if err != nil {
		log.Printf("could not copy favicon file with err; %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	profilePicUrl, err := vultr.UploadImage(profilePicFile, "user-profile-pictures")
	if err != nil {
		log.Printf("could not upload google user profile pic: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	args := sqlc.CreateUserParams{
		ID:              utils.RandomString(),
		Email:           req.Email,
		Name:            req.Email,
		EmailVerified:   true,
		Picture:         pgtype.Text{String: profilePicUrl, Valid: true},
		SignupMode:      sqlc.SignupModeEmail,
		AccountPassword: pgtype.Text{String: req.Password, Valid: true},
	}

	user, err := q.CreateUser(ctx, args)
	if err != nil {
		log.Printf("could not create user: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	issuedAt := time.Now().UTC()

	accessTokenDuration := config.Access_Token_Duration

	accessToken, _, err := token.CreateToken(user.ID, user.Name, user.Email, issuedAt, accessTokenDuration)
	if err != nil {
		log.Printf("could not create access token: %v", err)
		utils.Response(w, "could not create access token", http.StatusInternalServerError)
		return
	}

	refreshTokenDuration := config.Refresh_Token_Duration

	refreshToken, _, err := token.CreateToken(user.ID, user.Name, user.Email, issuedAt, refreshTokenDuration)
	if err != nil {
		log.Printf("could not create refresh token: %v", err)
		utils.Response(w, "could not create refresh token", http.StatusInternalServerError)
		return
	}

	updateUserLastLoginAndRefreshTokenParams := sqlc.UpdateUserLastLoginAndRefreshTokenParams{
		LastLogin:    pgtype.Timestamptz{Time: issuedAt, Valid: true},
		ID:           user.ID,
		RefreshToken: pgtype.Text{String: refreshToken, Valid: true},
	}

	user, err = q.UpdateUserLastLoginAndRefreshToken(ctx, updateUserLastLoginAndRefreshTokenParams)
	if err != nil {
		log.Printf("could not update user last login and refresh token:%v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	authenticatedUser := newAuthenticatedUser(&user, accessToken, refreshToken)

	if err := q.DeleteEmailVerificationToken(ctx, req.Email); err != nil {
		log.Printf("could not delete verification code: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	utils.JsonResponse(w, authenticatedUser)
}
