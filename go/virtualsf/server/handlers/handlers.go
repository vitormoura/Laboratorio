package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers/results"
	"net/http"
)

var (
	defaultStorageFactory model.VFStorageGroup
	defaultRouter         *mux.Router
)

func New(runInDebugMode bool, storageFactory model.VFStorageGroup) http.Handler {

	if defaultRouter != nil {
		return defaultRouter
	}

	defaultRouter := mux.NewRouter()
	defaultStorageFactory = storageFactory
	results.DebugMode = runInDebugMode

	handleVFolder(defaultRouter)
	handlePlayground(defaultRouter)
	handleCtrlPanel(defaultRouter)

	//Configuração do handler para servir conteúdo estático
	defaultRouter.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./server/static"))))

	return defaultRouter
}
