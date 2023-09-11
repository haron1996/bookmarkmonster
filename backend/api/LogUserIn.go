package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type logUseInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LogUserIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *logUseInReq

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

	user, err := q.GetUserByEmail(ctx, req.Email)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Printf("user with email: %s not found: %v", req.Email, err)
			utils.Response(w, fmt.Sprintf("user with email: %s not found", req.Email), http.StatusForbidden)
			return
		}
	}

	if !utils.CompareHash(req.Password, user.AccountPassword.String) {
		log.Println("could not log user in: wrong password")
		utils.Response(w, "wrong email or password", http.StatusForbidden)
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
