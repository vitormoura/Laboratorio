package storage

import (
	"encoding/json"
	"errors"
	_ "fmt"
	"github.com/pborman/uuid"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type vfdirStorage struct {
	root   string
	config storageConfig
}

func (dir *vfdirStorage) Add(f *model.File) error {

	if err := dir.verify(f); err != nil {
		return err
	}

	id := uuid.New()
	fileName := dir.getFileName(id)
	bytesGravados := int64(0)

	//Criando o arquivo principal
	file, err := os.Create(fileName)

	defer func() {
		file.Close()
	}()

	if err != nil {
		return err
	}

	if bytesGravados, err = io.Copy(file, f.Stream); err != nil {
		return err
	}

	f.ID = id
	f.Size = bytesGravados

	if bytesGravados == 0 {
		f.ID = ""
		os.Remove(fileName)
		return model.ErrEmptyFile
	}

	//Agora vamos criar o arquivo de metadados
	bytes, err := json.Marshal(f)

	if err != nil {
		os.Remove(fileName)
		f.ID = ""
		return err
	}

	mdfile, err := os.Create(fileName + ".meta")

	defer func() {
		mdfile.Close()
	}()

	if err != nil {
		f.ID = ""
		return err
	}

	mdfile.WriteString(string(bytes))

	return nil
}

func (dir *vfdirStorage) Find(id string) (*model.File, error) {
	//uniqID := uuid.Parse(id)

	fileName := dir.getMetaFileName(id)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil, model.ErrFileNotFound
	}

	//Recuperamos inicialmente a metadata do arquivo
	mdfileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var file model.File

	err = json.Unmarshal(mdfileBytes, &file)

	if err != nil {
		return nil, err
	}

	file.Stream, err = os.Open(dir.getFileName(id))

	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (dir *vfdirStorage) List() ([]model.FileInfo, error) {

	result := make([]model.FileInfo, 0, 20)

	filepath.Walk(dir.root, func(path string, info os.FileInfo, err error) error {

		if info != nil && !info.IsDir() {

			if filepath.Ext(info.Name()) == ".meta" {

				_, fileName := filepath.Split(info.Name())

				result = append(result, model.FileInfo{strings.Replace(fileName, ".file.meta", "", 1), ""})
			}

		} else if path != dir.root {

			//Diretórios diferentes da raiz não serão processados
			return filepath.SkipDir
		}

		return nil
	})

	return result, nil
}

func (dir *vfdirStorage) Stats() (model.VFStorageStats, error) {
	return model.VFStorageStats{}, nil
}

//verifyFilters testa se arquivo informado pode ser armazenado com base na configuração de filtros do diretório
func (dir *vfdirStorage) verify(file *model.File) error {

	if dir.config.Global.Locked {
		return model.ErrStorageLocked
	}

	if file.MimeType == "" {
		return model.ErrFileNotSupported
	}

	if len(dir.config.Filters.Allow) == 0 {
		return model.ErrFileNotSupported
	}

	mimeType := strings.ToLower(file.MimeType)

	//fmt.Println(file.MimeType)
	//fmt.Println(dir.config.Filters.Allow)
	//fmt.Println(sort.SearchStrings(dir.config.Filters.Allow, mimeType))

	if index := sort.SearchStrings(dir.config.Filters.Allow, mimeType); index < len(dir.config.Filters.Allow) && dir.config.Filters.Allow[index] == mimeType {
		return nil
	}

	return model.ErrFileNotSupported
}

//setup executa rotinas de inicializacao e configuração do mecanismo de armazenamento de arquivos locais,
//como criação do diretório e configuração padrão
func (dir *vfdirStorage) setup() error {

	//Vamos verificar se o diretório já existe:
	fi, err := os.Stat(dir.root)

	if err != nil && os.IsNotExist(err) {

		//Diretórios que não existem precisam ser criados e suas configurações inicializadas
		if err := os.Mkdir(dir.root, os.ModeDir); err != nil {
			return err
		}

		//Inicializando configuração básica
		dir.config = initConfigurationTo(dir.root)

	} else if !fi.IsDir() {
		return errors.New("Informe o caminho de um diretório válido")
	} else {
		dir.config = readConfigurationFrom(dir.root)
	}

	//ordenando filtros para execução de pesquisas posteriormente
	if len(dir.config.Filters.Allow) > 0 {
		sort.Strings(dir.config.Filters.Allow)
	}

	return nil
}

//getFileName recupera caminho completo do arquivo armazenado com ID informado
func (dir *vfdirStorage) getFileName(id string) string {
	return filepath.Join(dir.root, id+".file")
}

//getMetaFileName recupera caminho completo do arquivo de metadata do arquivo armazenado identificado pelo ID informado
func (dir *vfdirStorage) getMetaFileName(id string) string {
	return dir.getFileName(id) + ".meta"
}

//handleConfigurationUpdate atualiza configuração periodicamente (a cada minuto)
func (dir *vfdirStorage) handleConfigurationUpdate() {

	for _ = range time.Tick(1 * time.Minute) {
		dir.config = readConfigurationFrom(dir.root)
	}
}

//NewDirStorage cria um novo VFStorage que armazena arquivos em diretórios do
//sistema de arquivos a partir do diretório raiz informado. O modo longrunning indica
//que a nova instância será utilizada por muito tempo
func NewDirStorage(root string, longRunning bool) (model.VFStorage, error) {

	dir := vfdirStorage{root: root}
	err := dir.setup()

	if err != nil {
		return nil, err
	}

	if longRunning {
		go dir.handleConfigurationUpdate()
	}

	return &dir, nil
}
