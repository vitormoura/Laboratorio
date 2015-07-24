package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers/results"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/logs"
	"net/http"
)

var (
	defaultStorageFactory model.VFStorageGroup
	defaultRouter         *mux.Router
	serverlog             *logs.ServerLog
)

func New(runInDebugMode bool, storageFactory model.VFStorageGroup, srvlog *logs.ServerLog) http.Handler {

	if defaultRouter != nil {
		return defaultRouter
	}

	serverlog = srvlog
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
