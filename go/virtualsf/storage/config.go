package storage

import (
	"code.google.com/p/gcfg"
	"log"
	"os"
	"path/filepath"
)

type StorageConfig struct {
	StatsRefresh int
}

//localStorageConfig representa o esquema de configuração de uma pasta de armazenamento de arquivos
type localStorageConfig struct {
	Global struct {
		Locked bool
	}

	Filters struct {
		Allow []string
	}
}

//readConfigurationFrom lê arquivo de configuração do diretório e retorna sua representação em memória
func readConfigurationFrom(dir string) (localStorageConfig, error) {

	//Carregando configurações
	var (
		config localStorageConfig
		err    error
	)

	err = gcfg.ReadFileInto(&config, filepath.Join(dir, DIR_CONFIG_FILENAME))

	if err != nil {
		return config, err
	}

	return config, nil
}

//initConfigurationTo prepara um arquivo de configuração inicial para o diretório informado,
//retornando sua representação em memória
func initConfigurationTo(dir string) localStorageConfig {

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

	var config localStorageConfig

	gcfg.ReadStringInto(&config, configSample)

	return config
}
