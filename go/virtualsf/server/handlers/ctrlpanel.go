package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func handleCtrlPanel(r *mux.Router) {

	//Todas as actions vão exigir que o usuário seja o ADMIN
	r = r.PathPrefix("/ctrlpanel").MatcherFunc(onlyAdmin).Subrouter()

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		renderView("ctrlpanel/index", nil, w)
	})
}
