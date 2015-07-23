package results

import (
	"encoding/json"
	"fmt"
	"io"
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
	DebugMode bool

	//Cache de templates
	templateCache = make(map[string]*template.Template)
)

func View(viewName string, model interface{}, w http.ResponseWriter) {
	renderTemplate(viewName, model, w, false)
}

func Partial(viewName string, model interface{}, w http.ResponseWriter) {
	renderTemplate(viewName, model, w, true)
}

//content escreve o texto informado na saída enviada ao cliente
func Content(msg string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, msg)
}

//internalError retorna uma resposta do tipo erro ao cliente
func InternalError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)

	if DebugMode {
		fmt.Fprintln(w, err.Error())
	}
}

//notFound retorna um status do tipo 404, indicando que o recurso solicitado não foi encontrado
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

//writeFile escreve conteúdo binário do arquivo na saída http
func File(reader io.Reader, mimeType string, w http.ResponseWriter) {

	var (
		buffer        []byte
		qtdBytesLidos int
		err           error
	)

	w.Header().Set("Content-Type", mimeType)

	buffer = make([]byte, 10240)

	for qtdBytesLidos, err = reader.Read(buffer); qtdBytesLidos > 0; qtdBytesLidos, err = reader.Read(buffer) {

		if err != nil && err != io.EOF {
			InternalError(err, w)
			break
		}

		w.Write(buffer[0:qtdBytesLidos])
	}

}

//jsonr interpreta objeto model e retorna resposta do tipo json
func Json(model interface{}, w http.ResponseWriter) {

	bytes, err := json.Marshal(model)

	if err != nil {
		InternalError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(bytes))
}

func renderTemplate(viewName string, model interface{}, w http.ResponseWriter, isPartial bool) {

	var (
		t      *template.Template
		err    error
		exists bool
	)

	//Caso não exista o template no cache, vamos prepará-lo
	if t, exists = templateCache[viewName]; !exists || DebugMode {

		t = template.New("").Funcs(helpersFuncs)
		t, err = t.ParseFiles(mainLayoutFile, filepath.Join(templateDir, viewName+".html"))

		if err != nil {
			InternalError(err, w)
			return
		}

		templateCache[viewName] = t

		if DebugMode {
			log.Println(viewName, " compilada")
		}
	}

	w.Header().Set("Content-Type", "text/html")

	if !isPartial {
		err = t.ExecuteTemplate(w, "PAGE", model)
	} else {
		err = t.ExecuteTemplate(w, "MAIN", model)
	}

	if err != nil {
		InternalError(err, w)
	}
}
