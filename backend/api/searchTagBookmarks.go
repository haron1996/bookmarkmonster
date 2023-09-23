package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func SearchTagBookmarks(w http.ResponseWriter, r *http.Request) {
	tag_id := chi.URLParam(r, "tag_id")

	query := chi.URLParam(r, "query")

	ctx := r.Context()

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

	bookmarkTags, err := q.GetUserBookmarksByTagID(ctx, tag_id)
	if err != nil {
		log.Printf("could not get from bookmark tag: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	var results []sqlc.Bookmark

	for _, bt := range bookmarkTags {
		b, err := q.GetBookmarkByID(ctx, bt.BookmarkID)
		if err != nil {
			log.Printf("could not get bookmark by id: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		if strings.Contains(strings.ToLower(b.Title), strings.ToLower(query)) {
			results = append(results, b)
		}
	}

	utils.JsonResponse(w, results)
}
