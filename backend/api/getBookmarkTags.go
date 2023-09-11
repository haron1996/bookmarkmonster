package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func GetBookmarkTags(w http.ResponseWriter, r *http.Request) {
	bookmarkID := chi.URLParam(r, "bookmarkid")

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

	tagIDs, err := q.GetBookmarkTags(ctx, bookmarkID)
	if err != nil {
		log.Printf("could not get bookmark tags: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var tags []sqlc.Tag

	for _, tagID := range tagIDs {
		tag, err := q.GetTag(ctx, tagID)
		if err != nil {
			log.Printf("could not get tag: %v", err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		tags = append(tags, tag)
	}

	utils.JsonResponse(w, tags)
}
