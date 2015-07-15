package stats

import (
	_ "fmt"
	_ "github.com/vitormoura/Laboratorio/go/virtualsf/model"
	_ "io/ioutil"
	_ "os"
	_ "path/filepath"
	"time"
)

//DirStoreStatsAgent é um contínuo que atualiza estatísticas de utilização
//periodicamente
type DirStoreStatsAgent struct {
	root          string
	periodicidade time.Duration
	ticker        *time.Ticker
}

//Start inicia a execução do agente de atualização
func (d *DirStoreStatsAgent) Start() {

	if d.ticker != nil {
		return
	}

	d.ticker = time.NewTicker(d.periodicidade)

	go d.update()
}

//Stop interrompe a execução do agente de atualização
func (d *DirStoreStatsAgent) Stop() {
	if d.ticker != nil {
		d.ticker.Stop()
	}
}

func (d *DirStoreStatsAgent) update() {

	/*
		for agora := range d.ticker.C {

			filepath.Walk(d.root, func(path string, info os.FileInfo, err error) error {

				if info.IsDir() {

					stat := model.VFStorageStats{}

				}

				return nil
			})

		}
	*/
}

//NewAgent cria um novo agente que realiza atualizacoes a cada quantidade de minutos informados
func NewAgent(storageRoot string, minutes int) *DirStoreStatsAgent {

	return &DirStoreStatsAgent{root: storageRoot, periodicidade: time.Duration(minutes) * time.Minute}
}
