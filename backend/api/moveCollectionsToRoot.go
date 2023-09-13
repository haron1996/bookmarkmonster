package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type moveCollectionsToRootReq struct {
	Collections []sqlc.Folder `json:"collections"`
}

func MoveCollectionsToRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *moveCollectionsToRootReq

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

	var folders []sqlc.Folder

	for _, folder := range req.Collections {
		exists, err := q.GetRootFolderByName(ctx, folder.FolderName)
		if err != nil {
			log.Printf("could not get by name and null parent id: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		if exists {
			log.Println("root collections cannot share name")
			utils.Response(w, "root collections cannot share name", http.StatusForbidden)
			return
		}

		args := sqlc.MoveFoldersToRootParams{
			Label:   folder.Label,
			Label_2: folder.Label,
		}

		fs, err := q.MoveFoldersToRoot(ctx, args)
		if err != nil {
			log.Printf("folders not moved to root: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		if err := q.UpdateParentFolderToNull(ctx, folder.FolderID); err != nil {
			log.Printf("parent folder not updated to null: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		folders = append(folders, fs...)
	}

	utils.JsonResponse(w, folders)
}
