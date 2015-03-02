package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Home(r *mux.Router) {

	//home
	r.HandleFunc("/home", func(w http.ResponseWriter, req *http.Request) {
		view(w, "index", "Eu sou a mensagem exemplo")
	})
}
