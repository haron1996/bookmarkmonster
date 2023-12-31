package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func SearchBookmarks(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")

	const pLoad mw.ContextKey = "payload"

	ctx := r.Context()

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

	defer pool.Close()

	q := sqlc.New(pool)

	percent := "%"

	title = fmt.Sprintf("%s%s%s", percent, title, percent)

	bookmarks, err := q.SearchUserBookmarks(ctx, sqlc.SearchUserBookmarksParams{Title: title, UserID: payload.UserID})
	if err != nil {
		log.Printf("could not search bookmarks: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, bookmarks)
}
