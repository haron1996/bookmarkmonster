package mw

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/kwandapchumba/bookmarkmonster/db"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type ContextKey string

const pLoad ContextKey = "payload"

func AuthenticateRequest() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			payload, err := getAndVerifyToken(r)
			if err != nil {
				log.Printf("could not verify token: %v", err)
				utils.Response(w, "token is no longer valid", http.StatusUnauthorized)
				return
			}

			if payload != nil {

				account, err := db.GetUser(ctx, payload.UserEmail, payload.UserID)
				if err != nil {
					log.Println(err)
					utils.Response(w, "user does not exist", http.StatusUnauthorized)
					return
				}

				if payload.IssuedAt.Unix() != account.LastLogin.Time.Unix() {
					err := errors.New("token is no longer valid")
					log.Println(err)
					utils.Response(w, err.Error(), http.StatusUnauthorized)
					return
				}

				ctx := context.WithValue(r.Context(), pLoad, payload)

				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}

		return http.HandlerFunc(fn)
	}
}

func getAndVerifyToken(r *http.Request) (*token.PayLoad, error) {
	tkn := r.Header.Get("authorization")

	if tkn == "" {
		return nil, errors.New("token is empty")
	}

	splitToken := strings.Split(tkn, "Bearer")

	if len(splitToken) != 2 {
		return nil, errors.New("bearer token is not in proper format")
	}

	tkn = splitToken[1]

	payload, err := token.VerifyToken(tkn)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
