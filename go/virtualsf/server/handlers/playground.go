package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers/results"
	"net/http"
)

func handlePlayground(r *mux.Router) {
	r = r.PathPrefix("/playground").Subrouter()

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		results.View("playground/form", nil, w)
	})
}
