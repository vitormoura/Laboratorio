package server

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"io"
	_ "io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	_ "strings"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//FakeFile é uma implementação de io.Reader para permitir a
//simulação de envio de arquivos com bytes gerados aleatoriamente
type FakeFile struct {
	size    int
	currPos int
}

func (f *FakeFile) getBytesRemaining() int {
	return f.size - f.currPos
}

func (f *FakeFile) Read(p []byte) (n int, err error) {

	bytesRemaining := f.getBytesRemaining()
	bytesToRead := 0

	//fmt.Println(len(p), bytesRemaining)

	if bytesRemaining <= 0 {
		return 0, io.EOF
	}

	if len(p) == 0 {
		return 0, nil
	}

	if len(p) > bytesRemaining {
		bytesToRead = bytesRemaining
	} else {
		bytesToRead = len(p)
	}

	for i := 0; i < bytesToRead; i++ {
		p[i] = byte(letters[rand.Intn(len(letters))])
	}

	f.currPos += bytesToRead

	//fmt.Println(len(p), f.getBytesRemaining())

	return bytesToRead, nil
}

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

	if err != nil {
		return 0, "", err
	}

	//defer resp.Body.Close()
	//ioutil.ReadAll(resp.Body)

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

func TestSendLargeFile(t *testing.T) {

	fake := FakeFile{size: 1024 * 1024 * 40, currPos: 0} //40MB

	statusCode, fileID, err := sendFileToServer("myfile.txt", "text/plain", &fake)

	assert.Nil(t, err, "Requisição realizada sem erro")
	assert.Equal(t, 201, statusCode, "Server deve retornar código 201, indicando que um novo arquivo foi criado")
	assert.NotEqual(t, "", fileID, "O server deve retornar o ID do arquivo gerado através de um header")

}

func TestFakeFileReaderInterface(t *testing.T) {

	fake := FakeFile{size: 1024, currPos: 0}
	reader := bufio.NewReader(&fake)
	content, _ := reader.ReadString(' ')

	assert.Equal(t, 1024, len([]byte(content)), "Conteudo da string precisa ter 32000 caracteres")
}
