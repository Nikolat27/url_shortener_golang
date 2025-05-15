package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"url_shortener/internal/api"
)

func Routes(app *api.Application) *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/url/short/", app.UrlShorterHandler)

	return router
}
