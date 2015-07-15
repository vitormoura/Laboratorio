package main

import (
	"fmt"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage/stats"
)

func main() {
	config := server.GetDefaultConfiguration()

	agent := stats.NewAgent("c:\\", 1)
	agent.Start()

	fmt.Printf("-- HTTP SERVER LISTEN TO %d --\n", config.Main.ServerPort)
	server.Run(config)
}
