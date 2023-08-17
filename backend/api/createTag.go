package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type tag struct {
	Name string `json:"name"`
}

func (h *BaseHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *tag

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

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

	createTagParams := sqlc.CreateTagParams{
		ID:     uuid.New().String(),
		Name:   req.Name,
		UserID: payload.UserID,
	}

	tag, err := q.CreateTag(ctx, createTagParams)
	if err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "tag_user_id_name_key" (SQLSTATE 23505)` {
			utils.Response(w, "duplicate tag", http.StatusConflict)
			return
		}

		utils.Response(w, "could not create tag", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, tag)
}
