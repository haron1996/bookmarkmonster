package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type importFromPocketRequest struct {
	PocketAccessToken string `json:"pocket_access_token"`
	PocketUsername    string `json:"pocket_username"`
}

func ImportFromPocket(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *importFromPocketRequest

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

	const pLoad mw.ContextKey = "payload"

	ctx := r.Context()

	payload := ctx.Value(pLoad).(*token.PayLoad)

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Printf("could not load config: %v", err)
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

	token, err := q.GetPocketToken(ctx, payload.UserID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Printf("pocket token not found: %v", err)
			utils.Response(w, "pocket token not found", http.StatusNotFound)
			return
		default:
			log.Printf("could not get pocket token: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}
	}

	if req.PocketAccessToken != token.AccessToken || req.PocketUsername != token.Username || payload.UserID != token.UserID {
		log.Println("info mismatch")
		utils.Response(w, "something went wrong", 500)
		return
	}

	consumer_key := config.PocketConsumerKey
	access_token := token.AccessToken

	pocketAuthUrl := fmt.Sprintf("https://getpocket.com/v3/get?consumer_key=%s&access_token=%s&detailType=complete", consumer_key, access_token)

	request, err := http.NewRequest("POST", pocketAuthUrl, nil)
	if err != nil {
		log.Printf("could not create new request: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("X-Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("could not do request: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	var links map[string]interface{}

	err = json.Unmarshal(body, &links)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	var bookmarks []sqlc.Bookmark

	for _, item := range links["list"].(map[string]interface{}) {

		link := item.(map[string]interface{})

		host, url, err := getHost(fmt.Sprintf("%s", link["resolved_url"]))
		if err != nil {
			log.Printf("could not get url host: %s", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		title := ""

		resolved_title := fmt.Sprintf("%s", link["resolved_title"])

		if resolved_title == "" {
			title = fmt.Sprintf("%s", link["resolved_url"])
		} else {
			title = resolved_title
		}

		createBookmarkParams := sqlc.CreateBookmarkParams{
			ID:       utils.RandomString(),
			Title:    title,
			Bookmark: url,
			Host:     host,
			UserID:   payload.UserID,
			Notes:    pgtype.Text{},
			FolderID: pgtype.Text{},
		}

		bookmark, err := q.CreateBookmark(ctx, createBookmarkParams)
		if err != nil {
			log.Printf("could not create bookmark: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		bookmarks = append(bookmarks, bookmark)

		if tags, ok := link["tags"].(map[string]interface{}); ok {
			for tagName := range tags {
				getTagParams := sqlc.GetTagByUserAndNameParams{
					UserID: payload.UserID,
					Name:   strings.ToLower(tagName),
				}

				t, err := q.GetTagByUserAndName(ctx, getTagParams)
				if err != nil {
					switch {
					case errors.Is(err, pgx.ErrNoRows):

						ct, err := q.CreateTag(ctx, sqlc.CreateTagParams{ID: utils.RandomString(), Name: tagName, UserID: payload.UserID})
						if err != nil {
							log.Printf("could not create tag: %v", err)
							utils.Response(w, "something went wrong", 500)
							return
						}

						_, err = q.TagBookmark(ctx, sqlc.TagBookmarkParams{BookmarkID: bookmark.ID, TagID: ct.ID})
						if err != nil {
							log.Printf("could not tag bookmark: %v", err)
							utils.Response(w, "something went wrong", 500)
							return
						}

						continue

					default:
						log.Printf("could not get tag by user and name: %v", err)
						utils.Response(w, "something went wrong", 500)
						return
					}
				}

				_, err = q.TagBookmark(ctx, sqlc.TagBookmarkParams{BookmarkID: bookmark.ID, TagID: t.ID})
				if err != nil {
					log.Printf("could not tag bookmark: %v", err)
					utils.Response(w, "something went wrong", 500)
					return
				}

				continue
			}
		}

	}

	for _, bookmark := range bookmarks {
		params := sqlc.GetTagByUserAndNameParams{
			UserID: payload.UserID,
			Name:   strings.ToLower("pocket"),
		}

		pocketTag, err := q.GetTagByUserAndName(ctx, params)
		if err != nil {
			switch {
			case errors.Is(err, pgx.ErrNoRows):
				createTagParams := sqlc.CreateTagParams{
					ID:     utils.RandomString(),
					Name:   strings.ToLower("pocket"),
					UserID: payload.UserID,
				}

				tag, err := q.CreateTag(ctx, createTagParams)
				if err != nil {
					log.Printf("could not create tag: %v", err)
					utils.Response(w, "something went wrong", 500)
					return
				}

				tagBookmarkParams := sqlc.TagBookmarkParams{
					BookmarkID: bookmark.ID,
					TagID:      tag.ID,
				}

				_, err = q.TagBookmark(ctx, tagBookmarkParams)
				if err != nil {
					log.Printf("could not tag bookmark: %v", err)
					utils.Response(w, "something went wrong", 500)
					return
				}

				continue
			default:
				log.Printf("could not get pocket tag: %v", err)
				utils.Response(w, "something went wrong", 500)
				return
			}
		}

		tagBookmarkParams := sqlc.TagBookmarkParams{
			BookmarkID: bookmark.ID,
			TagID:      pocketTag.ID,
		}

		_, err = q.TagBookmark(ctx, tagBookmarkParams)
		if err != nil {
			log.Printf("could not tag bookmark: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}
	}

	utils.JsonResponse(w, bookmarks)
}

func getHost(URL string) (string, string, error) {
	if strings.Contains(URL, "?") {
		u, err := url.ParseRequestURI(URL)
		if err != nil {
			return "", "", err
		}

		if u.Scheme == "https" {

			return u.Host, fmt.Sprintf(`%v`, u), nil
		} else {
			return "", "", errors.New("invalid url")
		}

	} else {
		parsedUrl, err := url.Parse(URL)
		if err != nil {
			return "", "", err
		}

		if parsedUrl.Scheme == "https" {

			return parsedUrl.Host, URL, nil
		} else {

			return parsedUrl.String(), fmt.Sprintf(`https://%s`, URL), nil
		}

	}
}
