package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const host = "http://localhost:4040/vfolder/"

func TestInitializeServer(t *testing.T) {

	cmd, err := initServerDefaultConfiguration()

	assert.Nil(t, err, "Server inicializado sem erros")
	cmd.Process.Kill()
}

func TestSendValidSingleSmallFile(t *testing.T) {

	statusCode, fileID, err := sendFileToServer(host, "myfile.txt", "text/plain", bytes.NewBufferString("Eu sou um exemplo"))

	assert.Nil(t, err, "Requisição realizada sem erro")
	assert.Equal(t, 201, statusCode, "Server deve retornar código 201, indicando que um novo arquivo foi criado")
	assert.NotEqual(t, "", fileID, "O server deve retornar o ID do arquivo gerado através de um header")
}

func TestSendInvalidSingleSmallFile(t *testing.T) {

	statusCode, _, err := sendFileToServer(host, "myfile.pdf", "application/pdf", bytes.NewBufferString("Eu definitivamente não sou um arquivo PDF"))

	assert.Nil(t, err, "Requisição realizada sem erro")
	assert.Equal(t, 400, statusCode, "Server deve retornar código 400, não enviamos um arquivo com formato válido")
}

func TestGetFileList(t *testing.T) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", host+"files", nil)
	req.Header.Add("Authorization", "Basic "+getTestUserPassword())

	resp, err := client.Do(req)

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	assert.Nil(t, err, "Requisição realizada sem erro")
	assert.Equal(t, 200, resp.StatusCode, "Server deve retornar código 200")
	assert.Equal(t, "application/json", resp.Header["Content-Type"][0], "O header do tipo de retorno é application/json")

	var itens []interface{}
	err = decoder.Decode(&itens)

	assert.Nil(t, err, "Conteúdo da resposta é do tipo JSON")
	assert.NotNil(t, itens, "A resposta possui um objeto JSON válido")
}

func TestSendLargeFile(t *testing.T) {

	fake := NewFakeReader(1024 * 1024 * 40) //40MB

	statusCode, fileID, err := sendFileToServer(host, "myfile.txt", "text/plain", fake)

	assert.Nil(t, err, "Requisição realizada sem erro")
	assert.Equal(t, 201, statusCode, "Server deve retornar código 201, indicando que um novo arquivo foi criado")
	assert.NotEqual(t, "", fileID, "O server deve retornar o ID do arquivo gerado através de um header")
}
