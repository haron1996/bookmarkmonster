package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mailjet"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type requestEmailVerificationReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RequestEmailVerification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *requestEmailVerificationReq

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

	emailExists, err := q.CheckIfEmailIsInUse(ctx, req.Email)
	if err != nil {
		log.Printf("could not check if user exists: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	if emailExists {
		log.Println("email is already in use")
		utils.Response(w, "email is already in use", http.StatusForbidden)
		return
	}

	p, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Printf("could not hash password: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	args := sqlc.RequestEmailVerificationParams{
		Email:        req.Email,
		Code:         utils.GenerateEmailVerificationToken(),
		Expiry:       time.Now().UTC().Add(30 * time.Minute),
		UserPassword: p,
	}

	m := mailjet.NewEmailVerificationMail(args.Email, args.Code)

	if err := m.SendEmailVificationMail(); err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	v, err := q.RequestEmailVerification(ctx, args)
	if err != nil {
		log.Printf("could not request email verification: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	utils.JsonResponse(w, v)
}
