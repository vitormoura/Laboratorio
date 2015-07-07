package server

import (
	"code.google.com/p/gcfg"
	"crypto/sha1"
	"encoding/base64"
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

//GenerateSha1Password gera um password usando o algoritmo SHA-1 para ser utilizado na autenticação de usuários
func GenerateSha1Password(password string) string {

	data := []byte(password)
	d := sha1.New()
	d.Write(data)

	return string([]byte(base64.StdEncoding.EncodeToString(d.Sum(nil))))
}
