package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ControlPanel(r *mux.Router) {
	r = r.PathPrefix("/ctrlpanel").Subrouter()

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		renderView("ctrlpanel/index", nil, w)
	})
}
