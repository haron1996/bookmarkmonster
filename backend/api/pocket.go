package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type requestToken struct {
	Code  string
	State any
}

func RequestPocketAuthorizationCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Printf("could not load config: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	consumer_key := config.PocketConsumerKey
	redirect_uri := config.PocketRedirectUrl
	//state := ""
	pocketAuthUrl := fmt.Sprintf("https://getpocket.com/v3/oauth/request?consumer_key=%s&redirect_uri=%s", consumer_key, redirect_uri)

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

	var req requestToken

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

	authorizationUrl := fmt.Sprintf("https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s", req.Code, redirect_uri)

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		log.Printf("could not create new pool: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	defer pool.Close()

	q := sqlc.New(pool)

	_, err = q.CreatePocketCode(ctx, sqlc.CreatePocketCodeParams{UserID: payload.UserID, Code: req.Code})
	if err != nil {
		log.Printf("could not create pocket code: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	utils.JsonResponse(w, authorizationUrl)
}
