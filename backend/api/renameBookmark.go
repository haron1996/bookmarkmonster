package api

import (
	"encoding/json"
	"log"
	"net/http"

	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type renameBookmarkRequest struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

func (h *BaseHandler) RenameBookmark(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *renameBookmarkRequest

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

	bookmark, err := q.RenameBookmark(ctx, sqlc.RenameBookmarkParams{Title: req.Title, ID: req.ID})
	if err != nil {
		log.Printf("could not rename bookmark: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, bookmark)
}
