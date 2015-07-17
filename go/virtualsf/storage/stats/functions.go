package stats

import (
	"encoding/json"
	"fmt"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
	"os"
	"path"
	"path/filepath"
)

//calculateStatsFromDirStorageRoot preparar detalhes sobre o armazenamento de um storage com base em configurações padrão
func calculateStatsFromDirStorageRoot(dir string) (<-chan model.VFStorageStats, <-chan int, <-chan error) {

	resultC := make(chan model.VFStorageStats)
	errorC := make(chan error)
	doneC := make(chan int)

	go func() {

		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

			//Somente diretórios diferentes da raiz informada
			if info.IsDir() && path != dir {

				stat := model.VFStorageStats{App: info.Name(), Location: path}

				innerErr := filepath.Walk(path, func(innerPath string, innerInfo os.FileInfo, err error) error {

					if innerPath != path {

						if innerInfo.IsDir() {
							return filepath.SkipDir
						}

						if !storage.IsMetaFile(innerInfo.Name()) {
							stat.TotalSize += innerInfo.Size()
							stat.FileCount++
						}
					}

					return nil
				})

				if innerErr != nil {
					return innerErr
				}

				resultC <- stat

				//O conteúdo do diretório já foi avaliado internamente
				return filepath.SkipDir
			}

			return nil
		})

		if err != nil {
			errorC <- err
		}

		doneC <- 0

		close(resultC)
		close(errorC)
		close(doneC)

	}()

	return resultC, doneC, errorC
}

//saveStatsToDirStorage grava as estatísticas informadas no diretório padrão da pasta informada
func saveStatsToDirStorage(stats model.VFStorageStats) error {

	bytes, err := json.Marshal(stats)

	if err != nil {
		return err
	}

	statsFileName := fmt.Sprintf("stats-%d-%d-%d.json", stats.Date.Year(), stats.Date.Month(), stats.Date.Day())
	statsFile, err := os.Create(path.Join(stats.Location, storage.StatsDir, statsFileName))

	defer statsFile.Close()

	if err != nil {
		return err
	}

	statsFile.WriteString(string(bytes))

	return nil
}
