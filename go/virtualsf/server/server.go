package server

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/logs"
	"github.com/vitormoura/Laboratorio/go/virtualsf/services/refresher"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
	"log"
	"net/http"
	"time"
)

const (
	LOG_NAME                = "[server]"
	DEFAULT_STORAGE_REFRESH = 60
	DEFAULT_REQUEST_TIMEOUT = 10
)

//Run inicia execução do serviço de publicação e pesquisa de arquivos
func Run(config ServerConfig) {

	logger, err := logs.New(config.SharedFolder)

	if err != nil {
		log.Panic(err.Error())
	}

	storageFactory := storage.NewStorageFactory(config.SharedFolder)
	handler := handlers.New(config.DebugMode, config.TemplatesLocation, storageFactory, logger)

	//Servidor vai exigir autenticação do tipo BASIC com base em usuários e senhas do arquivo .htpasswd
	authenticator := auth.NewBasicAuthenticator("myrealm", auth.HtpasswdFileProvider(config.ServerUsersLocation))

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerPort),
		Handler:        auth.JustCheck(authenticator, handler.ServeHTTP),
		ReadTimeout:    DEFAULT_REQUEST_TIMEOUT * time.Second,
		WriteTimeout:   DEFAULT_REQUEST_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//Agente que atualiza estatísticas do storage
	agent := refresher.New(storageFactory, DEFAULT_STORAGE_REFRESH/60)
	agent.Start()

	log.Printf("%s iniciando servidor, escutando porta %d", LOG_NAME, config.ServerPort)

	go srv.ListenAndServe()
}
