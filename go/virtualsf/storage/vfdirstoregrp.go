package storage

import (
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"io"
	"os"
	"path/filepath"
)

type vfdirStorageGroup struct {
	root string
}

func (fac *vfdirStorageGroup) Get(appID string) (model.VFStorage, error) {
	return NewDirStorage(filepath.Join(fac.root, appID))
}

func (fac *vfdirStorageGroup) List() ([]string, error) {

	dir, err := os.Open(fac.root)

	if err != nil {
		return nil, err
	}

	result, err := dir.Readdirnames(0)

	if err != nil && err != io.EOF {
		return nil, err
	}

	return result, nil
}

func NewStorageFactory(rootDir string) model.VFStorageGroup {
	return &vfdirStorageGroup{rootDir}
}
