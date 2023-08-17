package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
	"github.com/kwandapchumba/bookmarkmonster/vultr"
)

type code struct {
	Code string `json:"code"`
}

type authenticatedUser struct {
	User         *sqlc.Userr
	AccessToken  string
	RefreshToken string
}

func newAuthenticatedUser(user *sqlc.Userr, accessToken, refreshToken string) *authenticatedUser {
	return &authenticatedUser{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (h *BaseHandler) RegisterGoogleUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var code *code

	if err := jsonDecoder.Decode(&code); err != nil {
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

	googleUser, err := utils.ExchangeGoogleAuthCodeForToken(code.Code, ctx)
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	q := sqlc.New(h.pool)

	getUserParams := sqlc.GetUserByEmailAndIDParams{
		Email: googleUser.Email,
		ID:    googleUser.Id,
	}

	var pgErr *pgconn.PgError

	userByGoogleEmailAndID, err := q.GetUserByEmailAndID(ctx, getUserParams)
	if err != nil {
		if errors.As(err, &pgErr) {
			log.Println(err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		} else if errors.Is(err, pgx.ErrNoRows) {
			// register user
			file, err := utils.DownloadGoogleUserProfilePic(googleUser.Picture)
			if err != nil {
				log.Println(err)
				utils.Response(w, "could not download profile pic", http.StatusInternalServerError)
				return
			}

			profilePicUrl, err := vultr.UploadUserProfilePic(file)
			if err != nil {
				log.Println(err)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}

			createUserParams := sqlc.CreateUserParams{
				ID:            googleUser.Id,
				Email:         googleUser.Email,
				Name:          googleUser.Name,
				EmailVerified: googleUser.EmailVerified,
				Picture:       pgtype.Text{String: profilePicUrl, Valid: true},
			}

			user, err := q.CreateUser(ctx, createUserParams)
			if err != nil {
				log.Println(err)
				utils.Response(w, "could not create new user", http.StatusInternalServerError)
				return
			}

			config, err := utils.LoadConfig(".")
			if err != nil {
				log.Println(err)
				utils.Response(w, "could not load config", http.StatusInternalServerError)
				return
			}

			issuedAt := time.Now().UTC()

			accessTokenDuration := config.Access_Token_Duration

			accessToken, _, err := token.CreateToken(user.ID, user.Name, user.Email, issuedAt, accessTokenDuration)
			if err != nil {
				log.Println(err)
				utils.Response(w, "could not create access token", http.StatusInternalServerError)
				return
			}

			refreshTokenDuration := config.Refresh_Token_Duration

			refreshToken, _, err := token.CreateToken(user.ID, user.Name, user.Email, issuedAt, refreshTokenDuration)
			if err != nil {
				log.Println(err)
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
				log.Println(err)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}

			authenticatedUser := newAuthenticatedUser(&user, accessToken, refreshToken)

			utils.JsonResponse(w, authenticatedUser)

			return
			// end register user
		} else {
			log.Println(err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	// login user
	loginUser(userByGoogleEmailAndID, w, ctx, q)
}

// utility to login user
func loginUser(user sqlc.Userr, w http.ResponseWriter, ctx context.Context, q *sqlc.Queries) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println(err)
		utils.Response(w, "could not load config", http.StatusInternalServerError)
		return
	}

	issuedAt := time.Now().UTC()

	accessTokenDuration := config.Access_Token_Duration

	accessToken, _, err := token.CreateToken(user.ID, user.Name, user.Email, issuedAt, accessTokenDuration)
	if err != nil {
		log.Println(err)
		utils.Response(w, "could not create access token", http.StatusInternalServerError)
		return
	}

	refreshTokenDuration := config.Refresh_Token_Duration

	refreshToken, _, err := token.CreateToken(user.ID, user.Name, user.Email, issuedAt, refreshTokenDuration)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	authenticatedUser := newAuthenticatedUser(&user, accessToken, refreshToken)

	utils.JsonResponse(w, authenticatedUser)
}
