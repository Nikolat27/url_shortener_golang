package errors

import (
	"net/http"
	"url_shortener/internal/helpers"
)

func ServerResponseError(w http.ResponseWriter, err error) {
	err = helpers.WriteJSON(w, err.Error(), http.StatusBadRequest)
	if err != nil {
		w.WriteHeader(500)
	}
}
