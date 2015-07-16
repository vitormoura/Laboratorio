package stats

import (
	"encoding/json"
	"fmt"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

//GetStatsFromDirStorage preparar detalhes sobre o armazenamento de um storage com base em configurações padrão
func getStatsFromDirStorage(dir string) (stat model.VFStorageStats, err error) {

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {

			files, err := ioutil.ReadDir(path)

			if err != nil {
				return err
			}

			for _, f := range files {

				//Arquivo não pode ser um diretório nem um META-FILE
				if !f.IsDir() && !storage.IsMetaFile(f.Name()) {
					stat.TotalSize += f.Size()
					stat.FileCount++
				}
			}

			return filepath.SkipDir
		}

		return nil
	})

	return
}

//saveStatsToDirStorage grava as estatísticas informadas no diretório padrão da pasta informada
func saveStatsToDirStorage(dir string, stats model.VFStorageStats) error {

	bytes, err := json.Marshal(stats)

	if err != nil {
		return err
	}

	statsFileName := fmt.Sprintf("stats-%d-%d-%d.json", stats.Date.Year(), stats.Date.Month(), stats.Date.Day())
	statsFile, err := os.Create(path.Join(dir, storage.StatsDir, statsFileName))

	defer statsFile.Close()

	if err != nil {
		return err
	}

	statsFile.WriteString(string(bytes))

	return nil
}
