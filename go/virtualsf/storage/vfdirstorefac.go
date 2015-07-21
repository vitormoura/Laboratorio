package storage

import (
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"os"
	"path/filepath"
)

type vfdirStorageFactory struct {
	root string
}

func (fac *vfdirStorageFactory) Create(appID string) (model.VFStorage, error) {
	return NewDirStorage(filepath.Join(fac.root, appID), false)
}

func (fac *vfdirStorageFactory) GetAvaiableStorages() ([]string, error) {

	result := make([]string, 0)

	filepath.Walk(fac.root, func(path string, info os.FileInfo, err error) error {

		//Somente diretórios diferentes da raiz informada
		if info.IsDir() && path != fac.root {

			result = append(result, info.Name())

			return filepath.SkipDir
		}

		return nil
	})

	return result, nil
}

func NewStorageFactory(rootDir string) model.VFStorageFactory {
	return &vfdirStorageFactory{rootDir}
}
