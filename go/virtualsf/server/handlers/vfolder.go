package handlers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
	"net/http"
	"time"
)

func VFolder(r *mux.Router) {
	r = r.PathPrefix("/vfolder").Subrouter()

	r.HandleFunc("/{app_id}", func(w http.ResponseWriter, req *http.Request) {

		var (
			app_id string
			exists bool
		)

		vars := mux.Vars(req)

		if app_id, exists = vars["app_id"]; !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		files, err := getFilesFromMultipartRequest(app_id, req)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err.Error())
			return
		}

		fs := storage.NewDirStorage("F:\\Temp\\virtualsf")

		for _, f := range files {
			fs.Add(&f)
		}

	})

	r.HandleFunc("/{app_id}/files/{id}", func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		id := vars["id"]
		fs := storage.NewDirStorage("D:\\Temp\\virtualsf")
		//app_id := vars["app_id"]

		file, err := fs.Find(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
			return
		}

		fmt.Fprintln(w, file)
	})

}

//extractFileFromMultipartRequest recupera objetos do tipo File lidos a partir da requisição informada
func getFilesFromMultipartRequest(appID string, req *http.Request) ([]model.File, error) {

	req.ParseMultipartForm(10240)

	if req.MultipartForm == nil {
		return nil, errors.New("Nenhum formulário do tipo multipart foi localizado")
	}

	if len(req.MultipartForm.File) == 0 {
		return nil, errors.New("Nenhum arquivo foi submetido")
	}

	//Vamos ler outros parâmetros que foram enviados no formulário e considerá-los propriedades do arquivo enviado
	properties := make(map[string]string)

	for k, v := range req.MultipartForm.Value {
		properties[k] = v[0]
	}

	files := make([]model.File, len(req.MultipartForm.File))
	pos := 0

	for _, f := range req.MultipartForm.File {

		for i := 0; i < len(f); i++ {

			//Preenchendo dados básicos do arquivo
			files[pos] = model.File{}
			files[pos].Name = f[i].Filename
			files[pos].PublishDate = time.Now()
			files[pos].App = appID
			files[pos].MimeType = f[i].Header["Content-Type"][0]
			files[pos].Properties = properties

			//Navegando para o fim do arquivo para descobrir seu tamanho completo
			if fileStream, err := f[i].Open(); err == nil {
				files[pos].Size, _ = fileStream.Seek(0, 2)
				files[pos].Stream = fileStream

				if files[pos].Size == 0 {
					return nil, errors.New("Arquivos vazios não serão publicados")
				}

				fileStream.Seek(0, 0)
			}
		}
	}

	return files, nil
}
