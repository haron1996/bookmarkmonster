package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func GetWaitList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := chi.URLParam(r, "code")

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	if code != config.AdminCode {
		log.Println("wrong code")
		utils.Response(w, "unauthorized", http.StatusForbidden)
		return
	}

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		log.Printf("could not create new pool: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	defer pool.Close()

	wl, err := sqlc.New(pool).GetWaitList(r.Context())
	if err != nil {
		log.Printf("could not get waitlist: %s", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	utils.JsonResponse(w, wl)
}
