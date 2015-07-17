package main

import (
	"github.com/vitormoura/Laboratorio/go/virtualsf/server"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
)

func main() {
	config := server.GetDefaultConfiguration()

	agent := storage.NewStatsUpdateAgent(config.Main.SharedFolder, 1)
	agent.Start()

	server.Run(config)
}
