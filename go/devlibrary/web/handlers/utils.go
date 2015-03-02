package handlers

import (
	"html/template"
	"log"
	"net/http"
	_ "os"
)

//var webWorkingDir string
var templateCache map[string]*template.Template

func init() {
	templateCache = make(map[string]*template.Template)
	//webWorkingDir, _ = os.Getwd()
}

//render renderiza o template localizado na pasta views
func view(w http.ResponseWriter, view string, model interface{}) {

	var (
		t   *template.Template
		ok  bool
		err error
	)

	if t, ok = templateCache[view]; !ok {
		t, err = template.ParseFiles("./templates/" + view + ".tpl")

		if err != nil {
			log.Printf("template not found : %s", view)
		}
	}

	if t != nil {
		t.Execute(w, model)
	} else {
		http.Error(w, "template not found", 500)
	}
}
