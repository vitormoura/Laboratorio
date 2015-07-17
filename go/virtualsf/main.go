package main

import (
	"github.com/vitormoura/Laboratorio/go/virtualsf/server"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
)

func main() {
	config := getDefaultConfiguration()

	agent := storage.NewStatsUpdateAgent(config.Server.SharedFolder, config.Storage.StatsRefresh)
	agent.Start()

	server.Run(config.Server)
}
