package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"net/http"
)

var (
	debugMode             bool
	defaultStorageFactory model.VFStorageFactory
	defaultRouter         *mux.Router
)

func New(debugMode bool, storageFactory model.VFStorageFactory) http.Handler {

	if defaultRouter != nil {
		return defaultRouter
	}

	defaultRouter := mux.NewRouter()
	defaultStorageFactory = storageFactory

	handleVFolder(defaultRouter)
	handlePlayground(defaultRouter)
	handleCtrlPanel(defaultRouter)

	//Configuração do handler para servir conteúdo estático
	defaultRouter.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./server/static"))))

	return defaultRouter
}
