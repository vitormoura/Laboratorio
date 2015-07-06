package server

import (
	"code.google.com/p/gcfg"
	"log"
)

//ServerConfig
type ServerConfig struct {
	Main struct {
		ServerPort   int
		SharedFolder string
		DebugMode    bool
	}
}

//getDefaultConfiguration recupera configurações do servidor com base em um arquivo server.ini
func GetDefaultConfiguration() ServerConfig {

	//Carregando configurações
	var (
		config ServerConfig
		err    error
	)

	err = gcfg.ReadFileInto(&config, "server.ini")

	if err != nil {
		log.Fatal("Configurações inválidas, não é possível iniciar o serviço : ", err.Error())
	}

	return config
}
