package tests

import (
	"encoding/base64"
	"fmt"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	"io"
	_ "log"
	"net/http"
)

func formatUserPassword(userName string, userPassword string) string {

	usrPlusPassword := fmt.Sprintf("%s:%s", userName, userPassword)
	usrPlusPassword = base64.StdEncoding.EncodeToString([]byte(usrPlusPassword))

	return usrPlusPassword
}

func configRequestAuth(req *http.Request, userName string, userPassword string) {
	authorization := "Basic " + formatUserPassword(userName, userPassword)
	req.Header.Add("Authorization", authorization)
}

func sendFileToServer(host string, userName string, userPassword string, fileName string, fileType string, fileContents io.Reader) (int, string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", host+fileName, fileContents)

	req.Header.Add("Content-Type", fileType)

	configRequestAuth(req, userName, userPassword)

	//log.Println("[POST]", host+fileName, authorization)
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
