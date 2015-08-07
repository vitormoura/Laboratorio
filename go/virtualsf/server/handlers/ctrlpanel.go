package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers/results"
)

func handleCtrlPanel(r *mux.Router) {

	//Todas as actions vão exigir que o usuário seja o ADMIN
	r = r.PathPrefix("/admin/ctrlpanel").Subrouter()

	//Action que exibe o painel de controle padrão
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		log.Println(LOG_NAME, "GET", req.URL.RequestURI())

		var (
			stats    map[string]*model.VFStorageStats
			appIDs   []string
			err      error
			currStat *model.VFStorageStats
		)

		stats = make(map[string]*model.VFStorageStats)
		appIDs, err = defaultStorageFactory.List()

		if err != nil {
			results.InternalError(err, w)
			return
		}

		for _, appID := range appIDs {
			grp, err := defaultStorageFactory.Get(appID)

			if err == nil {

				currStat, err = grp.Stats()

				if err == nil {
					stats[appID] = currStat
				}
			}
		}

		results.View("ctrlpanel/index", stats, w)
	})
}
