package server

import (
	_ "bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"io"
	"net/http"
	"os"
	"os/exec"
	_ "strings"
	"testing"
)

func initServerDefaultConfiguration() (*exec.Cmd, error) {
	os.Chdir("../")

	cmd := exec.Command("go", "run", "main.go")
	err := cmd.Start()

	return cmd, err
}

func getTestUserPassword() string {

	usrPlusPassword := fmt.Sprintf("%s:%s", "fula", "segredo")
	usrPlusPassword = base64.StdEncoding.EncodeToString([]byte(usrPlusPassword))

	return usrPlusPassword
}

func sendFileToServer(fileName string, fileType string, fileContents io.Reader) (int, string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:4040/vfolder/"+fileName, fileContents)

	req.Header.Add("Authorization", "Basic "+getTestUserPassword())
	req.Header.Add("Content-Type", fileType)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return 0, "", err
	}

	if fileID, exists := resp.Header[handlers.X_FILE_ID_HEADER]; exists {
		return resp.StatusCode, fileID[0], nil
	} else {
		return resp.StatusCode, "", nil
	}

}

func TestInitializeServer(t *testing.T) {

	cmd, err := initServerDefaultConfiguration()

	assert.Nil(t, err, "Server inicializado sem erros")
	cmd.Process.Kill()
}

func TestSendValidSingleSmallFile(t *testing.T) {
	//cmd, _ := initServerDefaultConfiguration()
	//defer cmd.Process.Kill()

	statusCode, fileID, err := sendFileToServer("myfile.txt", "text/plain", bytes.NewBufferString("Eu sou um exemplo"))

	assert.Nil(t, err, "Requisição realizada sem erro")
	assert.Equal(t, 201, statusCode, "Server deve retornar código 201, indicando que um novo arquivo foi criado")
	assert.NotEqual(t, "", fileID, "O server deve retornar o ID do arquivo gerado através de um header")
}

func TestSendInvalidSingleSmallFile(t *testing.T) {
	//cmd, _ := initServerDefaultConfiguration()
	//defer cmd.Process.Kill()

	statusCode, _, err := sendFileToServer("myfile.pdf", "application/pdf", bytes.NewBufferString("Eu definitivamente não sou um arquivo PDF"))

	assert.Nil(t, err, "Requisição realizada sem erro")
	assert.Equal(t, 400, statusCode, "Server deve retornar código 400, não enviamos um arquivo com formato válido")
}

func TestGetFileList(t *testing.T) {
	//cmd, _ := initServerDefaultConfiguration()
	//defer cmd.Process.Kill()

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:4040/vfolder/files", nil)
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