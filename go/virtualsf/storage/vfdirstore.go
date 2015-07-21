package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pborman/uuid"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type vfdirStorage struct {
	root   string
	config localStorageConfig
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

func (dir *vfdirStorage) Remove(id string) error {

	fileName := dir.getFileName(id)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return model.ErrFileNotFound
	}

	if err := os.Remove(fileName); err != nil {
		return err
	}

	metaFileName := dir.getMetaFileName(id)

	os.Remove(metaFileName)

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

			if isMetaFile(info.Name()) {

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

//Stats recupera estatisticas de armazenamento do storage
func (dir *vfdirStorage) Stats() (*model.VFStorageStats, error) {

	//Recuperamos inicialmente a metadata do arquivo
	statsFileBytes, err := ioutil.ReadFile(dir.getCurrentStatsFileName())

	if err != nil {
		return nil, err
	}

	var stats model.VFStorageStats

	err = json.Unmarshal(statsFileBytes, &stats)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}

//Refresh realiza rotinas de atualização do estado corrente do storage
func (dir *vfdirStorage) Refresh() error {

	stats, err := calcCurrentStats(dir.root)

	if err != nil {
		return err
	}

	return saveStatsToDirStorage(dir.root, stats)
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

		//Diretórios que não existem precisam ser criados e suas configurações padrão inicializadas

		//Diretório RAIZ
		if err := os.Mkdir(dir.root, os.ModeDir); err == nil {

			//Diretório LOGS
			if err := os.Mkdir(path.Join(dir.root, DIR_LOGS_LOCATION), os.ModeDir); err == nil {

				//Diretório STATS
				if err := os.Mkdir(path.Join(dir.root, DIR_STATS_LOCATION), os.ModeDir); err != nil {
					return err
				}

			} else {
				return err
			}

		} else {
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

//getMetaFileName recupera caminho completo do arquivo de metadata do arquivo armazenado identificado pelo ID informado
func (dir *vfdirStorage) getCurrentStatsFileName() string {
	return filepath.Join(dir.root, DIR_STATS_LOCATION, DIR_CURR_STATUS_FILENAME)
}

//handleConfigurationUpdate atualiza configuração periodicamente (a cada minuto)
func (dir *vfdirStorage) handleConfigurationUpdate() {

	for _ = range time.Tick(1 * time.Minute) {
		dir.config = readConfigurationFrom(dir.root)
	}
}

//isMetaFile verifica se o nome do arquivo informado o qualifica como um arquivo de metadados
func isMetaFile(fileName string) bool {
	return filepath.Ext(fileName) == ".meta"
}

//isConfigFile verifica se o nome do arquivo informado o qualifica como um arquivo de configuração
func isConfigFile(fileName string) bool {
	return filepath.Ext(fileName) == ".config"
}

func calcCurrentStats(dir string) (model.VFStorageStats, error) {

	stats := model.VFStorageStats{Date: time.Now()}

	innerErr := filepath.Walk(dir, func(innerPath string, innerInfo os.FileInfo, err error) error {

		if innerPath != dir {

			if innerInfo.IsDir() {
				return filepath.SkipDir
			}

			if !isMetaFile(innerInfo.Name()) && !isConfigFile(innerInfo.Name()) {
				stats.TotalSize += innerInfo.Size()
				stats.FileCount++
			}
		}

		return nil
	})

	if innerErr != nil {
		return stats, innerErr
	}

	return stats, nil
}

func saveStatsToDirStorage(dir string, stats model.VFStorageStats) error {

	bytes, err := json.Marshal(stats)

	if err != nil {
		return err
	}

	//Arquivo diário
	statsFileName := fmt.Sprintf("stats-%d-%d-%d.json", stats.Date.Year(), stats.Date.Month(), stats.Date.Day())
	statsFile, err := os.Create(path.Join(dir, DIR_STATS_LOCATION, statsFileName))

	defer statsFile.Close()

	if err != nil {
		return err
	}

	statsFile.WriteString(string(bytes))

	//Arquivo mais atual
	currentStatsFile, err := os.Create(path.Join(dir, DIR_STATS_LOCATION, DIR_CURR_STATUS_FILENAME))

	defer currentStatsFile.Close()

	if err != nil {
		return err
	}

	currentStatsFile.WriteString(string(bytes))

	return nil
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
