package main

import (
	"fmt"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server"
)

func main() {
	config := server.GetDefaultConfiguration()

	fmt.Printf("-- HTTP SERVER LISTEN TO %d --\n", config.Main.ServerPort)
	server.Run(config)
}
