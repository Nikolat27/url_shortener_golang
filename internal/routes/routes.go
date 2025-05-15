package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"url_shortener/internal/api"
)

func Routes(db *api.SqlDB) *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/url/short/", db.UrlShortHandler)
	return router
}
