package tests

import (
	"encoding/base64"
	"fmt"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func initServerDefaultConfiguration() (*exec.Cmd, error) {
	os.Chdir("../../")

	cmd := exec.Command("go", "run", "main.go config.go")
	err := cmd.Start()

	return cmd, err
}

func getTestUserPassword() string {

	usrPlusPassword := fmt.Sprintf("%s:%s", "fula", "segredo")
	usrPlusPassword = base64.StdEncoding.EncodeToString([]byte(usrPlusPassword))

	return usrPlusPassword
}

func sendFileToServer(host string, fileName string, fileType string, fileContents io.Reader) (int, string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", host+fileName, fileContents)

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
