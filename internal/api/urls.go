package api

import (
	"log/slog"
	"math/rand"
	"net/http"
	"url_shortener/internal/data"
	"url_shortener/internal/errors"
	"url_shortener/internal/helpers"
)

const (
	baseUrl = "http://localhost:8000/"
	characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func generateShortUrl(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}

func (app *Application) UrlShorterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		LongUrl string `json:"longUrl"`
		Password string `json:"password"`
		MaxUsage int `json:"max_usage"`
	}

	err := helpers.DeSerializeJSON(r.Body, 10000, &input)
	if err != nil {
		errors.ServerResponseError(w, err)
		return
	}

	shortUrl := generateShortUrl(8)

	url := &data.Url{
		LongUrl:  input.LongUrl,
		ShortUrl: shortUrl,
	}

	err = app.Models.Urls.Insert(url)
	if err != nil {
		errors.ServerResponseError(w, err)
		return
	}

	resp := baseUrl + url.ShortUrl
	err = helpers.WriteJSON(w, resp, http.StatusCreated)
	if err != nil {
		slog.Error("write json", "error", err)
	}
}

func (app *Application) RedirectLongUrlHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[1:]

	url, err := app.Models.Urls.Get(shortUrl)
	if err != nil {
		errors.ServerResponseError(w, err)
		return
	}

	http.Redirect(w, r, url.LongUrl, http.StatusPermanentRedirect)
}
