package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

//getAppID recupera identificador da aplicação da requisição
func getAppID(req *http.Request) string {
	return req.Header["X-Authenticated-Username"][0]
}

//onlyAdmin função do tipo filtro que só aceita requisições do usuário admin
func onlyAdmin(req *http.Request, m *mux.RouteMatch) bool {
	userID := getAppID(req)

	return strings.ToLower(userID) == "admin"
}
