package cookieapis

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func SearchTags(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("accessTokenCookie")
	if err != nil {
		log.Fatal(err)
	}

	payload, err := token.VerifyToken(c.Value)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(payload)

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		log.Printf("could not create new pool: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	pool.Config().MinConns = 50

	pool.Config().MaxConns = 150

	pool.Config().MaxConnIdleTime = 5 * time.Second

	defer pool.Close()

	q := sqlc.New(pool)

	tagName := chi.URLParam(r, "tagName")

	percent := "%"

	tagName = fmt.Sprintf("%s%s%s", percent, tagName, percent)

	params := sqlc.SearchTagsContainingTagNameAndUserIDParams{
		Name:   tagName,
		UserID: payload.UserID,
	}

	tags, err := q.SearchTagsContainingTagNameAndUserID(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tags)

	utils.JsonResponse(w, tags)
}
