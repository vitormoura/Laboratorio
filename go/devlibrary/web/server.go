package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/devlibrary/web/handlers"
	"net/http"
)

//Start inicia servidor web para escutar requisições na porta informada
func Start(port int) {
	r := mux.NewRouter()

	//Registrando handlers de requisição
	handlers.Home(r)

	//Registrando handler para servir conteúdo estático
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
