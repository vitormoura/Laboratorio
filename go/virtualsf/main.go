package main

import (
	"github.com/vitormoura/Laboratorio/go/virtualsf/server"
)

func main() {
	config := getDefaultConfiguration()
	server.Run(config.Server)
}
