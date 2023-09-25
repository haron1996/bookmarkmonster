package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func ConverPocketRequestTokenToAccessToken(w http.ResponseWriter, r *http.Request) {
	const pLoad mw.ContextKey = "payload"

	ctx := r.Context()

	payload := ctx.Value(pLoad).(*token.PayLoad)

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

	code, err := q.GetPocketCode(ctx, payload.UserID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Printf("pocket code does not exist: %s", err)
			utils.Response(w, "pocket code does not exist", http.StatusNotFound)
			return
		default:
			log.Printf("could not get pocket code: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}
	}

	consumer_key := config.PocketConsumerKey

	pocketAuthUrl := fmt.Sprintf("https://getpocket.com/v3/oauth/authorize?consumer_key=%s&code=%s", consumer_key, code.Code)

	request, err := http.NewRequest("POST", pocketAuthUrl, nil)
	if err != nil {
		log.Printf("could not create new request: %v", err)
		utils.Response(w, "something went wrong", http.StatusBadRequest)
		return
	}

	request.Header.Add("Content-Type", "application/json; charset=UTF-8")

	request.Header.Set("X-Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("could not do request: %v", err)
		utils.Response(w, "something went wrong", http.StatusBadRequest)
		return
	}

	defer response.Body.Close()

	if response.ContentLength == 0 {
		fmt.Println("Empty response from Pocket API")
		utils.Response(w, "something went wrong", http.StatusBadRequest)
		return
	}

	var req map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&req)
	if err != nil {
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

	access_token := req["access_token"]
	username := req["username"]

	token, err := q.CreatePocketToken(ctx, sqlc.CreatePocketTokenParams{AccessToken: fmt.Sprintf("%s", access_token), Username: fmt.Sprintf("%s", username), UserID: payload.UserID})
	if err != nil {
		log.Printf("could not create pocket token: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	utils.JsonResponse(w, token)
}
