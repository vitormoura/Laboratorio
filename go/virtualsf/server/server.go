package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/logs"
	"github.com/vitormoura/Laboratorio/go/virtualsf/services/refresher"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
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
	handler := handlers.New(config.DebugMode, config.TemplatesLocation, config.ServerUsersLocation, storageFactory, logger)

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerPort),
		Handler:        handler,
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
