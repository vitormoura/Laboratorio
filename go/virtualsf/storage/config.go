package storage

import (
	"code.google.com/p/gcfg"
	"log"
	"os"
	"path/filepath"
)

//storageConfig representa o esquema de configuração de uma pasta de armazenamento de arquivos
type storageConfig struct {
	Global struct {
		Locked bool
	}

	Filters struct {
		Allow []string
	}
}

//readConfigurationFrom lê arquivo de configuração do diretório e retorna sua representação em memória
func readConfigurationFrom(dir string) storageConfig {

	//Carregando configurações
	var (
		config storageConfig
		err    error
	)

	err = gcfg.ReadFileInto(&config, filepath.Join(dir, DIR_CONFIG_FILENAME))

	if err != nil {
		log.Fatal("Configurações inválidas, não é possível ler configurações do diretório : ", err.Error())
	}

	return config
}

//initConfigurationTo preparar um arquivo de configuração inicial para o diretório informado,
//retornando sua representação em memória
func initConfigurationTo(dir string) storageConfig {

	file, err := os.Create(filepath.Join(dir, DIR_CONFIG_FILENAME))

	defer func() {
		file.Close()
	}()

	if err != nil {
		log.Fatal(err.Error())
	}

	configSample := `
[Global]
locked=false

[Filters]
allow="text/plain"
allow="application/xml"
allow="text/xml"
allow="image/jpeg"
allow="image/png"
allow="image/bmp"
`

	file.WriteString(configSample)

	var config storageConfig

	gcfg.ReadStringInto(&config, configSample)

	return config
}
