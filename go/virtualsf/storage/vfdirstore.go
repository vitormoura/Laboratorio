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
	"path/filepath"
	"strings"
)

type vfdirStorage struct {
	root string
}

func (dir *vfdirStorage) Add(f *model.File) error {

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
		return errors.New("O arquivo nao possui ")
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

	fmt.Println(f.ID)

	return nil
}

func (dir *vfdirStorage) Find(id string) (*model.File, error) {
	//uniqID := uuid.Parse(id)

	fileName := dir.getMetaFileName(id)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil, nil
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
func NewDirStorage(root string) (model.VFStorage, error) {

	dir := vfdirStorage{root: root}
	fi, err := os.Stat(dir.root)

	if err != nil && os.IsNotExist(err) {

		if err := os.Mkdir(fi.Name(), os.ModeDir); err != nil {
			return nil, err
		}

	} else if !fi.IsDir() {
		return nil, errors.New("Informe o caminho de um diretório válido")
	}

	return &dir, nil
}
