package server

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"github.com/vitormoura/Laboratorio/go/virtualsf/services/refresher"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
	"log"
	"net/http"
	"time"
)

const LOG_NAME = "[server]"

//Run inicia execução do serviço de publicação e pesquisa de arquivos
func Run(config ServerConfig) {

	storageFactory := storage.NewStorageFactory(config.SharedFolder)
	handler := handlers.New(config.DebugMode, storageFactory)

	//Servidor vai exigir autenticação do tipo BASIC com base em usuários e senhas do arquivo .htpasswd
	authenticator := auth.NewBasicAuthenticator("myrealm", auth.HtpasswdFileProvider(".htpasswd"))

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerPort),
		Handler:        auth.JustCheck(authenticator, handler.ServeHTTP),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//Agente que atualiza estatísticas do storage
	agent := refresher.New(storageFactory, 1)
	agent.Start()

	log.Printf("%s iniciando servidor, escutando porta %d", LOG_NAME, config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}
