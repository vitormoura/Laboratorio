package refresher

import (
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"log"
	"time"
)

const LOG_NAME = "[services/refresher]"

//VFStorageRefresher é um agente contínuo que dispara a rotina de refresh dos storages registrados
//continuamente
type VFStorageRefresher struct {
	group         model.VFStorageGroup
	periodicidade time.Duration
	ticker        *time.Ticker
}

//Start inicia a execução do agente de atualização
func (d *VFStorageRefresher) Start() {

	if d.ticker != nil {
		return
	}

	d.ticker = time.NewTicker(d.periodicidade)

	go d.update()
}

//Stop interrompe a execução do agente de atualização
func (d *VFStorageRefresher) Stop() {
	if d.ticker != nil {
		d.ticker.Stop()
	}
}

func (d *VFStorageRefresher) update() {

	for agora := range d.ticker.C {

		log.Println(LOG_NAME, "efetuando refresh dos storages registrados", agora)

		storagesToRefresh, err := d.group.List()

		if err != nil {
			break
		}

		for _, appID := range storagesToRefresh {
			storage, err := d.group.Get(appID)

			if err != nil {
				continue
			}

			storage.Refresh()
		}

		log.Println(LOG_NAME, "refresh realizado")
	}
}

//New cria um novo agente que realiza atualizacoes a cada quantidade de minutos informados
func New(storageGroup model.VFStorageGroup, minutes int) *VFStorageRefresher {
	return &VFStorageRefresher{group: storageGroup, periodicidade: time.Duration(minutes) * time.Minute}
}
