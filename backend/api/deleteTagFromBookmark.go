package api

import (
	"encoding/json"
	"log"
	"net/http"

	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type deleteTagFromBookmarkRequest struct {
	BookmarkID string `json:"bookmarkid"`
	TagID      string `json:"tagid"`
}

func (h *BaseHandler) DeleteTagFromBookmark(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *deleteTagFromBookmarkRequest

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

	tagIDs, err := q.GetBookmarkTags(ctx, req.BookmarkID)
	if err != nil {
		log.Printf("could not get bookmark tags: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	if !(len(tagIDs) > 1) {

		log.Println("bookmark must have at least one tag")
		utils.Response(w, "bookmark must have at least one tag", http.StatusExpectationFailed)
		return
	}

	if err := q.DeleteTagFromBookmark(ctx, sqlc.DeleteTagFromBookmarkParams{BookmarkID: req.BookmarkID, TagID: req.TagID}); err != nil {
		log.Printf("could not delete tag from bookmark: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	tag, err := q.GetTag(ctx, req.TagID)
	if err != nil {
		log.Printf("could nto get tag: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, tag)
}
