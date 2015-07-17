package stats

import (
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"log"
	"time"
)

const logName = "[storage/stats]"

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

	for agora := range d.ticker.C {

		log.Println(logName, "iniciando atualizacao de estatisticas")

		var (
			stats model.VFStorageStats
			err   error
		)

		resultC, doneC, errorC := calculateStatsFromDirStorageRoot(d.root)

	STATS_CALC_LOOP:
		for {
			select {
			case stats = <-resultC:

				stats.Date = agora
				err = saveStatsToDirStorage(stats)

				if err != nil {
					break
				}

			case err = <-errorC:
				log.Println(logName, err.Error())

			case <-doneC:
				log.Println(logName, "estatisticas atualizadas")
				break STATS_CALC_LOOP
			}
		}
	}
}

//NewAgent cria um novo agente que realiza atualizacoes a cada quantidade de minutos informados
func NewAgent(storageRoot string, minutes int) *DirStoreStatsAgent {

	return &DirStoreStatsAgent{root: storageRoot, periodicidade: time.Duration(minutes) * time.Minute}
}
