package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers/results"
)

const (

	//X_FILE_ID_HEADER é o header HTTP enviado contendo o ID de um arquivo recém criado
	X_FILE_ID_HEADER   string = "X-File-Id"
	X_FILE_NAME_HEADER string = "X-File-Name"
	LOG_NAME                  = "[server/handlers]"
)

func handleVFolder(r *mux.Router) {

	r = r.PathPrefix("/api/vfolder").Subrouter()

	get := r.Methods("GET").Subrouter()
	post := r.Methods("POST").Subrouter()

	//Action para publicar um novo arquivo através de um formulário de envio de arquivos tradicional
	post.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		log.Println(LOG_NAME, "POST", req.URL.RequestURI())

		var (
			files []*model.File
			err   error
			appID string
		)

		appID = getAppID(req)
		files, err = getFilesFromMultipartRequest(appID, req)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err.Error())
			return
		}

		createFiles(appID, files, w)

		results.View("vfolder/created", files, w)

		serverlog.Info(appID, LOG_NAME, "NEW FILE")
	})

	//Action para publicar um novo arquivo com base no corpo da requisição
	post.HandleFunc("/{file_name}", func(w http.ResponseWriter, req *http.Request) {

		log.Println(LOG_NAME, "POST", req.URL.RequestURI())

		var (
			vars            map[string]string
			files           []*model.File
			err             error
			appID, fileName string
		)

		vars = mux.Vars(req)
		appID = getAppID(req)
		fileName = vars["file_name"]

		files, err = getFilesFromRESTRequest(appID, fileName, req)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err.Error())
			return
		}

		createFiles(appID, files, w)

		serverlog.Info(appID, LOG_NAME, "NEW FILE")
	})

	//Action para listar em formato JSON uma lista de dados básicos dos arquivos de uma determinada aplicação
	get.HandleFunc("/files", func(w http.ResponseWriter, req *http.Request) {

		log.Println(LOG_NAME, "GET ", req.URL.RequestURI())

		var (
			err   error
			files []model.FileInfo
			fs    model.VFStorage
			appID string
		)

		appID = getAppID(req)
		fs, err = getDefaultStorage(appID)

		if err != nil {
			results.InternalError(err, w)
			return
		}

		files, err = fs.List()

		if err != nil {
			results.InternalError(err, w)
			return
		}

		results.Json(files, w)

		serverlog.Info(appID, LOG_NAME, "GET FILE LIST")
	})

	//Action para remover um arquivo
	r.Methods("DELETE").Subrouter().HandleFunc("/files/{id}", func(w http.ResponseWriter, req *http.Request) {

		log.Println(LOG_NAME, "DELETE ", req.URL.RequestURI())

		var (
			vars  map[string]string
			id    string
			err   error
			fs    model.VFStorage
			appID string
		)

		vars = mux.Vars(req)
		id = vars["id"]
		appID = getAppID(req)
		fs, err = getDefaultStorage(appID)

		if err != nil {
			results.InternalError(err, w)
			return
		}

		err = fs.Remove(id)

		if err != nil {

			if err == model.ErrFileNotFound {
				w.WriteHeader(http.StatusNotFound)
			} else {
				results.InternalError(err, w)
			}

			return
		}

		w.WriteHeader(http.StatusOK)

		serverlog.Info(appID, LOG_NAME, "REMOVE FILE")
	})

	//Action para realizar o download do arquivo identificado pelo ID informado
	get.HandleFunc("/files/{id}", func(w http.ResponseWriter, req *http.Request) {
		log.Println(LOG_NAME, "GET ", req.URL.RequestURI())

		var (
			vars map[string]string
			id   string
			err  error
			file *model.File
			fs   model.VFStorage
		)

		vars = mux.Vars(req)
		id = vars["id"]
		fs, err = getDefaultStorage(getAppID(req))

		if err != nil {
			results.InternalError(err, w)
			return
		}

		file, err = fs.Find(id)

		if err != nil {

			//O erro retornado pode apontar que o arquivo não existe para o usuário
			if err == model.ErrFileNotFound {
				results.NotFound(w)
				return
			}

			results.InternalError(err, w)
			return
		}

		//Inclui headers com detalhes sobre o arquivo
		w.Header().Add(X_FILE_ID_HEADER, file.ID)
		w.Header().Add(X_FILE_NAME_HEADER, file.Name)

		results.File(file.Stream, file.MimeType, w)

		serverlog.Info(file.App, LOG_NAME, "GET FILE")
	})

	//Action para recuperar situação das estatísticas de armazenamento da aplicação
	get.HandleFunc("/stats/stats.json", func(w http.ResponseWriter, req *http.Request) {
		log.Println(LOG_NAME, "GET ", req.URL.RequestURI())

		appID := getAppID(req)
		fs, err := getDefaultStorage(appID)

		if err != nil {
			results.InternalError(err, w)
			return
		}

		stats, err := fs.Stats()

		if err != nil {
			results.InternalError(err, w)
			return
		}

		if stats == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		results.Json(stats, w)

		serverlog.Info(appID, LOG_NAME, "GET JSON STATS")
	})
}

//createFiles grava arquivos junto ao sistema de armazenamento e formata resposta para os clientes
func createFiles(appID string, files []*model.File, w http.ResponseWriter) {

	fs, err := getDefaultStorage(appID)

	if err != nil {
		results.InternalError(err, w)
		return
	}

	for _, f := range files {

		if err := fs.Add(f); err != nil {

			if err == model.ErrFileNotSupported {
				w.WriteHeader(http.StatusBadRequest)
			} else if err == model.ErrStorageLocked {
				w.WriteHeader(http.StatusForbidden)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			fmt.Fprintln(w, err.Error())
			return
		}

		w.Header().Add(X_FILE_ID_HEADER, f.ID)
		w.Header().Add("Location", "/vfolder/files/"+f.ID)
	}

	w.WriteHeader(http.StatusCreated)
}

//getDefaultStorage recupera storage padrão para arquivos
func getDefaultStorage(appID string) (model.VFStorage, error) {
	return defaultStorageFactory.Get(appID)
}

//getFilesFromMultipartRequest recupera objetos do tipo File lidos a partir da requisição informada
func getFilesFromMultipartRequest(appID string, req *http.Request) ([]*model.File, error) {

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

	files := make([]*model.File, len(req.MultipartForm.File))
	pos := 0

	for _, f := range req.MultipartForm.File {

		for i := 0; i < len(f); i++ {

			//Preenchendo dados básicos do arquivo
			_, fileName := filepath.Split(f[i].Filename)

			files[pos] = new(model.File)
			files[pos].Name = fileName
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
func getFilesFromRESTRequest(appID string, fileName string, req *http.Request) ([]*model.File, error) {

	files := make([]*model.File, 1)
	pos := 0

	//Preenchendo dados básicos do arquivo
	files[pos] = new(model.File)
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
