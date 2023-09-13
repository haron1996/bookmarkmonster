package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type moveCollectionToAnotherReq struct {
	DestinationCollection sqlc.Folder   `json:"destination_folder"`
	FoldersToMove         []sqlc.Folder `json:"folders_to_move"`
}

func MoveCollectionsToAnother(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *moveCollectionToAnotherReq

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

	for _, collectionToMove := range req.FoldersToMove {
		args := sqlc.GetFolderByNameAndParentIDParams{
			FolderName:  collectionToMove.FolderName,
			SubfolderOf: pgtype.Text{String: req.DestinationCollection.FolderID, Valid: true},
		}

		exists, err := q.GetFolderByNameAndParentID(ctx, args)
		if err != nil {
			log.Printf("could not get foler by name and parent id: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		if exists {
			log.Println("collection must not have more that one child with the same name")
			utils.Response(w, "collection must not have more that one children with the same name", http.StatusForbidden)
			return
		}

		args2 := sqlc.MoveCollectionsParams{
			Label:   req.DestinationCollection.Label,
			Label_2: collectionToMove.Label,
			Label_3: collectionToMove.Label,
		}

		fs, err := q.MoveCollections(ctx, args2)
		if err != nil {
			log.Printf("could not move collections: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		args3 := sqlc.UpdateFolderParentParams{
			SubfolderOf: pgtype.Text{String: req.DestinationCollection.FolderID, Valid: true},
			FolderID:    collectionToMove.FolderID,
		}

		if err := q.UpdateFolderParent(ctx, args3); err != nil {
			log.Printf("could not update folder parent: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		folders = append(folders, fs...)
	}

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

	for _, f := range folders {
		args := sqlc.GetFolderParams{
			FolderID: f.FolderID,
			UserID:   payload.UserID,
		}

		folder, err := q.GetFolder(ctx, args)
		if err != nil {
			log.Printf("could not get folder: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		folders = append(folders, folder)
	}

	utils.JsonResponse(w, folders)
}
