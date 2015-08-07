package handlers

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers/results"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/logs"

	auth "github.com/abbot/go-http-auth"
)

var (
	defaultStorageFactory    model.VFStorageGroup
	defaultRouter            *mux.Router
	serverlog                *logs.ServerLog
	defaultUserAuthenticator auth.AuthenticatorInterface
)

func New(runInDebugMode bool, templatesLocation string, usersPasswordFileLocation string, storageFactory model.VFStorageGroup, srvlog *logs.ServerLog) http.Handler {

	if defaultRouter != nil {
		return defaultRouter
	}

	serverlog = srvlog
	defaultUserAuthenticator = auth.NewBasicAuthenticator("myrealm", auth.HtpasswdFileProvider(usersPasswordFileLocation))
	defaultStorageFactory = storageFactory
	results.TemplatesDir = templatesLocation
	results.DebugMode = runInDebugMode

	defaultRouter = mux.NewRouter()
	adminRouter := mux.NewRouter()
	apiRouter := mux.NewRouter()

	handleVFolder(apiRouter)
	handleCtrlPanel(adminRouter)
	handlePlayground(apiRouter)

	defaultRouter.PathPrefix("/api").Handler(
		negroni.New(
			negroni.HandlerFunc(authorized),
			negroni.Wrap(apiRouter),
		),
	)

	defaultRouter.PathPrefix("/admin").Handler(
		negroni.New(
			negroni.HandlerFunc(authorized),
			negroni.HandlerFunc(adminOnly),
			negroni.Wrap(adminRouter),
		),
	)

	defaultRouter.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./server/static"))))

	return defaultRouter
}
