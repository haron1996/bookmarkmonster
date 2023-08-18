package api

import (
	"encoding/json"
	"log"
	"net/http"

	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type tagBookmarkRequest struct {
	BookmarkID string     `json:"bookmark_id"`
	Tags       []sqlc.Tag `json:"tags"`
}

func (h *BaseHandler) TagBookmark(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *tagBookmarkRequest

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

	// const pLoad mw.ContextKey = "payload"

	// payload := ctx.Value(pLoad).(*token.PayLoad)

	q := sqlc.New(h.pool)

	for _, tag := range req.Tags {
		tagBookmarkParams := sqlc.TagBookmarkParams{
			BookmarkID: req.BookmarkID,
			TagID:      tag.ID,
		}
		_, err := q.TagBookmark(ctx, tagBookmarkParams)
		if err != nil {
			log.Printf("could not tag bookmark: %v", err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	utils.Response(w, "bookmark tagged successfully", http.StatusOK)
}
