package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

const (
	templateDir    = "./server/templates"
	mainLayoutFile = templateDir + "/layout.html"
)

var (

	//Cache de templates
	templateCache = make(map[string]*template.Template)

	//DebugMode determina se as funções relacionadas ao processamento de requisições devem atuar em modo debug
	DebugMode bool = false
)

//view escreve o resultado do processamento do template indicado, usando dados
func renderView(viewName string, model interface{}, w http.ResponseWriter) {

	var (
		t      *template.Template
		err    error
		exists bool
	)

	//Caso não exista o template no cache, vamos prepará-lo
	if t, exists = templateCache[viewName]; !exists || DebugMode {

		t, err = template.ParseFiles(mainLayoutFile, filepath.Join(templateDir, viewName+".html"))

		if err != nil {
			internalError(w)

			if DebugMode {
				fmt.Fprintln(w, err.Error())
			}

			return
		}

		templateCache[viewName] = t

		if DebugMode {
			log.Println(viewName, " compilada")
		}
	}

	w.Header().Set("Content-Type", "text/html")

	err = t.Execute(w, model)

	if err != nil {
		log.Println("Erro ao processar template ", viewName)
		internalError(w)
	}
}

//content escreve o texto informado na saída enviada ao cliente
func content(msg string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, msg)
}

//internalError retorna uma resposta do tipo erro ao cliente
func internalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

//notFound retorna um status do tipo 404, indicando que o recurso solicitado não foi encontrado
func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

//jsonr interpreta objeto model e retorna resposta do tipo json
func jsonr(model interface{}, w http.ResponseWriter) {

	bytes, err := json.Marshal(model)

	if err != nil {
		internalError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(bytes))
}
