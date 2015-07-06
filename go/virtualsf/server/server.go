package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"log"
	"net/http"
	"time"
)

//Run inicia execução do serviço de publicação e pesquisa de arquivos
func Run(config ServerConfig) {

	r := mux.NewRouter()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", r)

	handlers.DebugMode = config.Main.DebugMode
	handlers.VFolder(r, config.Main.SharedFolder)
	handlers.Playground(r)

	srv := &http.Server{

		Addr:           fmt.Sprintf(":%d", config.Main.ServerPort),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServe())
}
