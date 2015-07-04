package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Playground(r *mux.Router) {
	r = r.PathPrefix("/playground").Subrouter()

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		renderView("playground/form", nil, w)
	})
}
