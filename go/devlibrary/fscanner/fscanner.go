package fscanner

import (
	"crypto/md5"
	"fmt"
	"github.com/vitormoura/Laboratorio/go/devlibrary/bookrp"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type extFileInfo struct {
	info os.FileInfo
	path string
}

func computeMd5FromFile(filePath string) ([]byte, error) {

	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	return hash.Sum(result), nil
}

func findAllDirectoryFiles(dir string) chan extFileInfo {

	newFileCh := make(chan extFileInfo)
	dirContents, err := ioutil.ReadDir(dir)

	if err == nil && len(dirContents) > 0 {

		go func() {

			for _, el := range dirContents {

				fullName := filepath.Join(dir, el.Name())

				if el.IsDir() {

					for i := range findAllDirectoryFiles(fullName) {
						newFileCh <- i
					}

				} else {
					newFileCh <- extFileInfo{info: el, path: fullName}
				}
			}

			close(newFileCh)

		}()

	} else {
		close(newFileCh)
	}

	return newFileCh
}

//ScanBooksFromDir recupera livros para todos os arquivos a partir do diretório raiz informado
func ScanBooksFromDir(dir string) chan bookrp.Book {

	newBookFoundCh := make(chan bookrp.Book)

	go func() {
		for f := range findAllDirectoryFiles(dir) {

			id, _ := computeMd5FromFile(f.path)

			newBookFoundCh <- bookrp.Book{Id: fmt.Sprintf("%x", id), Name: f.info.Name(), Size: f.info.Size(), Location: f.path}
		}

		close(newBookFoundCh)
	}()

	return newBookFoundCh
}
