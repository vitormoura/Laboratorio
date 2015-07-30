package main

import (
	"log"

	"code.google.com/p/gcfg"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server"
	"github.com/vitormoura/Laboratorio/go/virtualsf/storage"
)

type GlobalConfig struct {
	Server  server.ServerConfig
	Storage storage.StorageConfig
}

//getDefaultConfiguration recupera configurações do servidor com base em um arquivo server.ini
func getDefaultConfiguration() GlobalConfig {

	//Carregando configurações
	var (
		config GlobalConfig
		err    error
	)

	err = gcfg.ReadFileInto(&config, "virtualsf.ini")

	if err != nil {
		log.Fatal("Configurações inválidas, não é possível iniciar o serviço : ", err.Error())
	}

	return config
}
