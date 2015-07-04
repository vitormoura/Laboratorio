package storage

import (
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type vfdirStorage struct {
	root string
}

func (dir *vfdirStorage) Add(f *model.File) error {

	id := uuid.New()
	fileName := dir.getFileName(id)

	//Criando o arquivo principal
	file, err := os.Create(fileName)

	defer func() {
		file.Close()
	}()

	if err != nil {
		return err
	}

	if _, err := io.Copy(file, f.Stream); err != nil {
		return err
	}

	f.ID = id

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

	//Recuperamos inicialmente a metadata do arquivo
	mdfileBytes, err := ioutil.ReadFile(dir.getMetaFileName(id))

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

		fmt.Println(path)

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

func (dir *vfdirStorage) getFileName(id string) string {
	return filepath.Join(dir.root, id+".file")
}

func (dir *vfdirStorage) getMetaFileName(id string) string {
	return dir.getFileName(id) + ".meta"
}

//NewDirStorage cria um novo VFStorage que armazena arquivos em diretórios do
//sistema de arquivos a partir do diretório raiz informado
func NewDirStorage(root string) model.VFStorage {

	dir := vfdirStorage{root: root}

	return &dir
}
