package server

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"log"
	"net/http"
	"time"
)

const LOG_NAME = "[server]"

//Run inicia execução do serviço de publicação e pesquisa de arquivos
func Run(config ServerConfig) {

	port := config.ServerPort
	router := mux.NewRouter()

	//Configurações do handlers
	handlers.DebugMode = config.DebugMode

	http.Handle("/", router)

	handlers.VFolder(router, config.SharedFolder)
	handlers.Playground(router)
	handlers.ControlPanel(router)

	//Configuração do handler para servir conteúdo estático
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./server/static"))))

	//Servidor vai exigir autenticação do tipo BASIC com base em usuários e senhas do arquivo .htpasswd
	authenticator := auth.NewBasicAuthenticator("myrealm", auth.HtpasswdFileProvider(".htpasswd"))

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        auth.JustCheck(authenticator, router.ServeHTTP),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("%s iniciando servidor, escutando porta %d", LOG_NAME, port)
	log.Fatal(srv.ListenAndServe())
}
