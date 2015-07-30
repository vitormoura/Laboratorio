package server

import (
	"crypto/sha1"
	"encoding/base64"
)

//ServerConfig
type ServerConfig struct {
	ServerPort          int
	DebugMode           bool
	ServerUsersLocation string
	SharedFolder        string
	TemplatesLocation   string
}

//GenerateSha1Password gera um password usando o algoritmo SHA-1 para ser utilizado na autenticação de usuários
func GenerateSha1Password(password string) string {

	if password == "" {
		return password
	}

	data := []byte(password)
	d := sha1.New()
	d.Write(data)

	return string([]byte(base64.StdEncoding.EncodeToString(d.Sum(nil))))
}
