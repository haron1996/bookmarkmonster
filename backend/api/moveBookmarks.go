package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type moveBookmarksReq struct {
	Bookmarks   []sqlc.Bookmark `json:"bookmarks"`
	Destination sqlc.Folder     `json:"destination"`
}

func MoveBookmarks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *moveBookmarksReq

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

	var bookmarks []sqlc.Bookmark

	for _, bookmark := range req.Bookmarks {
		var args sqlc.MoveBookmarkParams

		if req.Destination.FolderID == "" {
			args = sqlc.MoveBookmarkParams{
				FolderID: pgtype.Text{},
				ID:       bookmark.ID,
			}
		} else {
			args = sqlc.MoveBookmarkParams{
				FolderID: pgtype.Text{String: req.Destination.FolderID, Valid: true},
				ID:       bookmark.ID,
			}
		}

		b, err := q.MoveBookmark(ctx, args)
		if err != nil {
			log.Printf("could not move bookmarks: %s", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		bookmarks = append(bookmarks, b)
	}

	utils.JsonResponse(w, bookmarks)
}
