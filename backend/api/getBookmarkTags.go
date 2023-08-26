package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func (h *BaseHandler) GetBookmarkTags(w http.ResponseWriter, r *http.Request) {
	bookmarkID := chi.URLParam(r, "bookmarkid")

	ctx := r.Context()

	q := sqlc.New(h.pool)

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
