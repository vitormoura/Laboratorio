package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers/results"
)

func handlePlayground(r *mux.Router) {
	r = r.PathPrefix("/api/playground").Subrouter()

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		results.View("playground/form", nil, w)
	})
}
