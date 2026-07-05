package handler

import (
	"net/http"

	"github.com/athulanilthomas/www/api/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router, err := router.BuildRouter()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	router.ServeHTTP(w, r)
}
