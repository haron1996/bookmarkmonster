package api

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type addBookmarkRequest struct {
	Bookmark string     `json:"bookmark"`
	Tags     []sqlc.Tag `json:"tags"`
}

func (h *BaseHandler) AddBookmark(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *addBookmarkRequest

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

	var host string

	var urlToOpen string

	if strings.Contains(req.Bookmark, "?") {
		u, err := url.ParseRequestURI(req.Bookmark)
		if err != nil {
			log.Println(err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		if u.Scheme == "https" {

			host = u.Host

			urlToOpen = fmt.Sprintf(`%v`, u)
		} else {
			utils.Response(w, "invalid url", http.StatusBadRequest)
			return
		}

	} else {
		parsedUrl, err := url.Parse(req.Bookmark)
		if err != nil {
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		if parsedUrl.Scheme == "https" {
			host = parsedUrl.Host

			urlToOpen = req.Bookmark
		} else {
			host = parsedUrl.String()

			urlToOpen = fmt.Sprintf(`https://%s`, req.Bookmark)
		}

	}

	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	page := browser.MustPage(urlToOpen).MustWaitStable()

	defer browser.MustClose()

	body := page.MustElement("body")

	body.MustEval(`() => this.style.overflowY = 'hidden'`)

	page.MustScreenshot("bookmarkScreenshot.png")

	pageHasTitle, title, err := page.Has("title")
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var pageTitle string

	if pageHasTitle {
		pageTitle = strings.TrimSpace(title.MustText())
	}

	var bookmarkFaviconURL string
	var faviconLocationEmpty bool

	resp, err := http.Get(fmt.Sprintf("https://www.google.com/s2/favicons?domain=%v&sz=64", urlToOpen))
	if err != nil {
		log.Printf("error getting favicon url: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	log.Printf("response statuscode is: %v", resp.StatusCode)

	if resp.StatusCode != 200 {
		log.Println("favicon location empty")
		faviconLocationEmpty = true
		bookmarkFaviconURL = ""
	}

	log.Printf("bookmark favicon location is empty: %v", faviconLocationEmpty)

	if !faviconLocationEmpty {
		faviconLocation := resp.Header.Get("content-location")

		log.Printf("favicon location is: %v", faviconLocation)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS12,
			},
		}

		client := &http.Client{Transport: tr}

		request, err := http.NewRequest("GET", faviconLocation, nil)
		if err != nil {
			log.Printf("error getting favicon: %v", err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		// request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0")
		// request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

		response, err := client.Do(request)
		if err != nil {
			log.Printf("could not send request: %v", err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		defer response.Body.Close()

		// if response.StatusCode != 200 {
		// 	log.Printf("coudl not get favicon with status code: %v", response.StatusCode)
		// 	// utils.Response(w, "something went wrong", http.StatusInternalServerError)
		// 	// return
		// 	bookmarkFaviconURL = ""
		// }

		if response.StatusCode == 200 {
			file, err := os.Create("bookmarkFavicon.png")
			if err != nil {
				log.Println(err)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}

			defer file.Close()

			_, err = io.Copy(file, response.Body)
			if err != nil {
				log.Println(err)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}
		} else {
			faviconLocationEmpty = true
		}
	}

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	reg := config.VultrRegion

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(config.VultrAccessKey, config.VultrSecretKey, ""),
		Endpoint:         aws.String(fmt.Sprintf("https://%s", config.VultrHostname)),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String(reg),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	s3Client := s3.New(newSession)

	// load bookmark screenshot
	bookmarkScreenshotFile, err := os.Open("bookmarkScreenshot.png")
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	// upload bookmark screenshot
	s3InputBookmarkScreenshotObject := s3.PutObjectInput{
		Bucket: aws.String("/bookmark-screenshots"),
		Key:    aws.String(uuid.New().String()),
		Body:   bookmarkScreenshotFile,
		ACL:    aws.String("public-read"),
	}

	_, err = s3Client.PutObject(&s3InputBookmarkScreenshotObject)
	if err != nil {
		if err != nil {
			log.Println(err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	bookmarkScreenshotURL := fmt.Sprintf("https://%s.vultrobjects.com/bookmark-screenshots/%s", reg, *s3InputBookmarkScreenshotObject.Key)

	if !faviconLocationEmpty {
		// load bookmark favicon
		bookmarkFaviconFile, err := os.Open("bookmarkFavicon.png")
		if err != nil {
			log.Println(err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		// upload bookmark favicon
		s3InputBookmarkFaviconObject := s3.PutObjectInput{
			Bucket: aws.String("/bookmark-favicons"),
			Key:    aws.String(uuid.New().String()),
			Body:   bookmarkFaviconFile,
			ACL:    aws.String("public-read"),
		}

		_, err = s3Client.PutObject(&s3InputBookmarkFaviconObject)
		if err != nil {
			if err != nil {
				log.Println(err)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}
		}

		bookmarkFaviconURL = fmt.Sprintf("https://%s.vultrobjects.com/bookmark-favicons/%s", reg, *s3InputBookmarkFaviconObject.Key)
	}

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

	q := sqlc.New(h.pool)

	params := sqlc.AddBookmarkParams{
		ID:        uuid.New().String(),
		Title:     pageTitle,
		Bookmark:  urlToOpen,
		Host:      host,
		Favicon:   bookmarkFaviconURL,
		Thumbnail: bookmarkScreenshotURL,
		UserID:    payload.UserID,
	}

	bookmark, err := q.AddBookmark(ctx, params)
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	// tag bookmark
	for _, tag := range req.Tags {
		params := sqlc.BookmarkTagParams{
			BookmarkID: bookmark.ID,
			TagID:      tag.ID,
		}

		_, err := q.BookmarkTag(ctx, params)
		if err != nil {
			var pgErr *pgconn.PgError
			switch {
			case errors.As(err, &pgErr):
				log.Printf("could not create bookmark: %v", pgErr)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			default:
				log.Printf("could not create bookmark: %v", err)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}
		}
	}

	utils.JsonResponse(w, bookmark)
}
