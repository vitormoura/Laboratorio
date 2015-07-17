package main

import (
	"github.com/vitormoura/Laboratorio/go/virtualsf/server"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage/stats"
)

func main() {
	config := server.GetDefaultConfiguration()

	agent := stats.NewAgent("D:\\Temp\\virtualsf-tests\\", 1)
	agent.Start()

	server.Run(config)
}
