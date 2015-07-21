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
	factory       model.VFStorageFactory
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

		storagesToRefresh, err := d.factory.GetAvaiableStorages()

		if err != nil {
			break
		}

		for _, appID := range storagesToRefresh {
			storage, err := d.factory.Create(appID)

			if err != nil {
				continue
			}

			storage.Refresh()
		}

		log.Println(LOG_NAME, "refresh realizado")
	}
}

//New cria um novo agente que realiza atualizacoes a cada quantidade de minutos informados
func New(storageFactory model.VFStorageFactory, minutes int) *VFStorageRefresher {
	return &VFStorageRefresher{factory: storageFactory, periodicidade: time.Duration(minutes) * time.Minute}
}
