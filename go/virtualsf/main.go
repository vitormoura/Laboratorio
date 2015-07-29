package main

import (
	"fmt"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server"
	"os"
	"os/signal"
)

func main() {
	config := getDefaultConfiguration()

	server.Run(config.Server)

	//Observando sinais do OS para encerrar execução
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	<-c

	fmt.Println("Bye bye !")
}
