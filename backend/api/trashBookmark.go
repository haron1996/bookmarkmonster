package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type trashBookmarkRequest struct {
	Bookmarks []sqlc.Bookmark `json:"bookmarks"`
}

func (h *BaseHandler) TrashBookmark(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *trashBookmarkRequest

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

	q := sqlc.New(h.pool)

	var bookmarks []sqlc.Bookmark

	for _, b := range req.Bookmarks {
		bookmark, err := q.TrashBookmark(ctx, sqlc.TrashBookmarkParams{Deleted: pgtype.Timestamptz{Time: time.Now().UTC(), Valid: true}, ID: b.ID})
		if err != nil {
			log.Printf("could not trash bookmark: %v", err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		bookmarks = append(bookmarks, bookmark)
	}

	utils.JsonResponse(w, bookmarks)
}
