package api

import (
	"log/slog"
	"math/rand"
	"net/http"
	"url_shortener/internal/data"
	"url_shortener/internal/errors"
	"url_shortener/internal/helpers"
)

func generateShortUrl(length int) string {
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}

func (app *Application) UrlShorterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		LongUrl string `json:"longUrl"`
	}

	err := helpers.DeSerializeJSON(r.Body, 10000, &input)
	if err != nil {
		errors.ServerResponseError(w, err)
		return
	}

	shortUrl := generateShortUrl(8)

	url := &data.Url{
		LongUrl: input.LongUrl,
		ShortUrl: shortUrl,
	}
	
	err = app.Models.Urls.Insert(url)
	if err != nil {
		errors.ServerResponseError(w, err)
		return
	}
	
	err = helpers.WriteJSON(w, "Done!", 200)
	if err != nil {
		slog.Error("write json", "error", err)
	}
}
