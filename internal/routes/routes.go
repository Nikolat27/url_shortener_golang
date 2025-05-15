package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"url_shortener/internal/api"
)

func Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/url/short/", api.UrlShortHandler)
	return router
}
