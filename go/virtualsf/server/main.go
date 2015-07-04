package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", r)

	handlers.DebugMode = true
	handlers.VFolder(r)
	handlers.Playground(r)

	fmt.Println("-- HTTP SERVER LISTEN TO 4040 --")
	http.ListenAndServe(":4040", nil)
}
