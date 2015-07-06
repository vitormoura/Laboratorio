package handlers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
	"net/http"
	_ "strconv"
	"time"
)

const (

	//X_FILE_ID_HEADER é o header HTTP enviado contendo o ID de um arquivo recém criado
	X_FILE_ID_HEADER string = "X-FILE-ID"
)

var (
	folderStorageLocation string
)

func VFolder(r *mux.Router, sharedFolder string) {

	folderStorageLocation = sharedFolder

	r = r.PathPrefix("/vfolder/{app_id}").Subrouter()

	get := r.Methods("GET").Subrouter()
	post := r.Methods("POST").Subrouter()

	//Action para publicar um novo arquivo através de um formulário de envio de arquivos tradicional
	post.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		var (
			vars  map[string]string
			files []model.File
			err   error
		)

		vars = mux.Vars(req)
		files, err = getFilesFromMultipartRequest(vars["app_id"], req)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err.Error())
			return
		}

		createFiles(files, w)
	})

	//Action para publicar um novo arquivo com base no corpo da requisição
	post.HandleFunc("/{file_name}", func(w http.ResponseWriter, req *http.Request) {

		if req.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var (
			vars            map[string]string
			files           []model.File
			err             error
			appID, fileName string
		)

		vars = mux.Vars(req)
		appID = vars["app_id"]
		fileName = vars["file_name"]

		files, err = getFilesFromRESTRequest(appID, fileName, req)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err.Error())
			return
		}

		createFiles(files, w)

	})

	//Action para listar em formato JSON uma lista de dados básicos dos arquivos de uma determinada aplicação
	get.HandleFunc("/files", func(w http.ResponseWriter, req *http.Request) {

		var (
			err   error
			files []model.FileInfo
			fs    model.VFStorage
		)

		fs = getDefaultStorage()
		files, err = fs.List()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
			return
		}

		jsonr(files, w)
	})

	//Action para realizar o download do arquivo identificado pelo ID informado
	get.HandleFunc("/files/{id}", func(w http.ResponseWriter, req *http.Request) {

		var (
			vars map[string]string
			id   string
			err  error
			file *model.File
			fs   model.VFStorage
		)

		vars = mux.Vars(req)
		id = vars["id"]
		fs = getDefaultStorage()

		file, err = fs.Find(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
			return
		}

		if file == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch req.Method {

		case "GET":
			writeFile(file.Stream, file.MimeType, w)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})

}

//createFiles grava arquivos junto ao sistema de armazenamento e formata resposta para os clientes
func createFiles(files []model.File, w http.ResponseWriter) {

	fs := getDefaultStorage()

	for _, f := range files {

		if err := fs.Add(&f); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
			return
		}

		w.Header().Add(X_FILE_ID_HEADER, f.ID)
	}

	w.WriteHeader(http.StatusCreated)
}

//getDefaultStorage recupera storage padrão para arquivos
func getDefaultStorage() model.VFStorage {
	return storage.NewDirStorage(folderStorageLocation)
}

//getFilesFromMultipartRequest recupera objetos do tipo File lidos a partir da requisição informada
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

//getFilesFromMultipartRequest recupera objetos do tipo File lidos a partir do corpo da requisiçao
func getFilesFromRESTRequest(appID string, fileName string, req *http.Request) ([]model.File, error) {

	files := make([]model.File, 1)
	pos := 0

	//Preenchendo dados básicos do arquivo
	files[pos] = model.File{}
	files[pos].Name = fileName
	files[pos].PublishDate = time.Now()
	files[pos].App = appID

	if len(req.Header["Content-Type"]) == 0 {
		return nil, errors.New("Formato do conteúdo do arquivo não foi informado")
	}

	if req.ContentLength == 0 {
		return nil, errors.New("Nenhum dado enviado na requisição")
	}

	files[pos].MimeType = req.Header["Content-Type"][0]
	files[pos].Stream = req.Body

	return files, nil
}
