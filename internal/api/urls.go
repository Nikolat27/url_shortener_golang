package api

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"url_shortener/internal/errors"
	"url_shortener/internal/helpers"
)

type Url struct {
	LongUrl string `json:"longUrl"`
}

type SqlDB struct {
	DB *sql.DB	
}

func generateShortUrl(length int) string {
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}

func (s SqlDB) UrlShortHandler(w http.ResponseWriter, r *http.Request) {
	var url Url

	err := helpers.DeSerializeJSON(r.Body, 10000, &url)
	if err != nil {
		errors.ServerResponseError(w, err)
		return
	}

	shortUrl := generateShortUrl(8)
	fmt.Println(shortUrl)
}
