package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func SearchUserTags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tagName := chi.URLParam(r, "tagName")

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

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

	pool.Config().MinConns = 50

	pool.Config().MaxConns = 150

	pool.Config().MaxConnIdleTime = 5 * time.Second

	defer pool.Close()

	q := sqlc.New(pool)

	percent := "%"

	tagName = fmt.Sprintf("%s%s%s", percent, tagName, percent)

	params := sqlc.SearchTagsContainingTagNameAndUserIDParams{
		Name:   tagName,
		UserID: payload.UserID,
	}

	tags, err := q.SearchTagsContainingTagNameAndUserID(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError

		switch {
		case errors.As(err, &pgErr):
			log.Println(err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		case errors.Is(err, pgx.ErrNoRows):
			log.Println(err)
			utils.Response(w, "no tags found", http.StatusNoContent)
			return
		default:
			log.Println(err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	utils.JsonResponse(w, tags)
}
