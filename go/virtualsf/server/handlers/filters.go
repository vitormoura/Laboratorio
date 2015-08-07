package handlers

import (
	"net/http"
	"strings"

	"github.com/abbot/go-http-auth"
)

//getAppID recupera identificador da aplicação da requisição
func getAppID(req *http.Request) string {

	if val, exists := req.Header["X-Authenticated-Username"]; exists {
		return val[0]
	} else {
		return ""
	}
}

func adminOnly(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	userID := getAppID(req)

	if strings.ToLower(userID) == "admin" {
		next(w, req)
	} else {
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

func authorized(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	auth.JustCheck(defaultUserAuthenticator, next)(w, r)
}
