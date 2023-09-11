package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mailjet"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type resendConfirmationLinkRequest struct {
	Email string `json:"email"`
}

func ResendConfirmationLink(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *resendConfirmationLinkRequest

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

	args := sqlc.UpdateEmailVericationTokenParams{
		Expiry: time.Now().UTC().Add(30 * time.Minute),
		Email:  link.Email,
	}

	if err := q.UpdateEmailVericationToken(ctx, args); err != nil {
		log.Printf("could not update email verification token: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	m := mailjet.NewEmailVerificationMail(link.Email, link.Code)

	if err := m.SendEmailVificationMail(); err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	msg := "confirmation link has been resent"

	utils.JsonResponse(w, msg)
}
