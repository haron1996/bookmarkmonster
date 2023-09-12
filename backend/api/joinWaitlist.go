package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type joinWaitlistReq struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CompanyName string `json:"company_name"`
	Comment     string `json:"comment"`
}

func JoinWaitList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *joinWaitlistReq

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

	args := sqlc.JoinWaitListParams{
		ID:          utils.RandomString(),
		Email:       req.Email,
		Name:        req.Name,
		CompanyName: pgtype.Text{String: req.CompanyName, Valid: true},
		Comment:     pgtype.Text{String: req.Comment, Valid: true},
	}

	wl, err := q.JoinWaitList(ctx, args)
	if err != nil {
		log.Printf("could not join waitlist: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	utils.JsonResponse(w, wl)
}
