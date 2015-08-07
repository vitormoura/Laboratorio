package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	_ "log"
	"net/http"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
)

const (
	defaultHost         = "http://localhost:4045"
	defaultVFolderHost  = defaultHost + "/api/vfolder/"
	defaultUserName     = "temp"
	defaultServerPort   = 4045
	defaultUserPassword = "segredo"
)

func TestServer(t *testing.T) {

	startServer() //

	Convey("PUBLICAÇÃO DE ARQUIVOS", t, func() {

		Convey("Arquivos Simples", func() {

			Convey("usuários não identificados não podem publicar arquivos", func() {
				statusCode, fileID, err := sendFileToServer(defaultVFolderHost, "EU_NAO_EXISTO", defaultUserPassword, "myfile.txt", "text/plain", bytes.NewBufferString("Eu sou um exemplo"))

				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 401)
				So(fileID, ShouldBeEmpty)
			})

			Convey("Arquivo Pequeno Formato Permitido retorna sucesso", func() {

				statusCode, fileID, err := sendFileToServer(defaultVFolderHost, defaultUserName, defaultUserPassword, "myfile.txt", "text/plain", bytes.NewBufferString("Eu sou um exemplo"))

				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 201)
				So(fileID, ShouldNotBeEmpty)
			})

			Convey("Arquivo Pequeno Formato Não Permitido não é publicado", func() {

				statusCode, fileID, err := sendFileToServer(defaultVFolderHost, defaultUserName, defaultUserPassword, "myfile.pdf", "application/pdf", bytes.NewBufferString("DUMMY DATA"))

				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 400)
				So(fileID, ShouldBeEmpty)
			})

			Convey("Exclusao arquivo publicado retorna sucesso", func() {

				//Publicamos um arquivo
				statusCode, fileID, err := sendFileToServer(defaultVFolderHost, defaultUserName, defaultUserPassword, "myfile.txt", "text/plain", bytes.NewBufferString("DUMMY DATA"))

				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 201)
				So(fileID, ShouldNotBeEmpty)

				//Em seguida o excluímos
				client := &http.Client{}

				req, err := http.NewRequest("DELETE", defaultVFolderHost+"files/"+fileID, nil)
				configRequestAuth(req, defaultUserName, defaultUserPassword)

				resp, err := client.Do(req)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, 200)
			})

			Convey("Exclusao arquivo não existente retorna falha", func() {

				client := &http.Client{}

				req, err := http.NewRequest("DELETE", defaultVFolderHost+"files/7931cd97-6611-4e69-b3d4-dcbac4dba3c6", nil)
				configRequestAuth(req, defaultUserName, defaultUserPassword)

				resp, err := client.Do(req)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, 404)
			})

			Convey("Arquivo Grande Formato Permitido retorna sucesso", func() {

				fakeFile := NewFakeReader(1024 * 1024 * 40) //40MB
				statusCode, fileID, err := sendFileToServer(defaultVFolderHost, defaultUserName, defaultUserPassword, "myfile.txt", "text/plain", fakeFile)

				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 201)
				So(fileID, ShouldNotBeEmpty)
			})

			Convey("repositório do usuário está bloqueado, não é possível publicar novos arquivos", func() {

				statusCode, fileID, err := sendFileToServer(defaultVFolderHost, "APP_3", defaultUserPassword, "myfile.txt", "text/plain", bytes.NewBufferString("Eu sou um exemplo"))

				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 403)
				So(fileID, ShouldBeEmpty)

			})

			Convey("recuperar lista de arquivos recupera coleção json", func() {

				client := &http.Client{}

				req, err := http.NewRequest("GET", defaultVFolderHost+"files", nil)
				configRequestAuth(req, "APP_2", defaultUserPassword)

				resp, err := client.Do(req)

				defer resp.Body.Close()

				So(resp.StatusCode, ShouldEqual, 200)
				So(resp, ShouldNotBeNil)
				So(err, ShouldBeNil)

				var result []model.FileInfo

				json.NewDecoder(resp.Body).Decode(&result)

				So(len(result), ShouldEqual, 3)

			})

		})

		SkipConvey("Múltiplos Arquivos", func() {

			Convey("Envio Arquivos Concorrentemente grava todos corretamente", func() {

				files := map[string]int{
					"myfile1.txt":  20,
					"myfile2.txt":  5,
					"myfile3.txt":  3,
					"myfile4.txt":  35,
					"myfile5.txt":  12,
					"myfile6.txt":  1,
					"myfile7.txt":  47,
					"myfile8.txt":  6,
					"myfile9.txt":  5,
					"myfile10.txt": 3,
					"myfile11.txt": 12,
					"myfile12.txt": 1,
					"myfile13.txt": 47,
				}

				resc := make(chan bool)
				allSucceed := true

				for k, v := range files {

					go func(fileName string, fileSize int) {
						fakeFile := NewFakeReader(1024 * 1024 * fileSize)

						statusCode, fileID, err := sendFileToServer(defaultVFolderHost, defaultUserName, defaultUserPassword, fileName, "text/plain", fakeFile)

						resc <- (err == nil && statusCode == 201 && fileID != "")
					}(k, v)
				}

				for i := 0; i < len(files); i++ {
					select {
					case res := <-resc:
						if !res {
							allSucceed = false
						}
					}
				}

				So(allSucceed, ShouldBeTrue)

			})

		})

		clearTempFiles()
	})

	Convey("GERAÇÃO DE SENHAS (SHA)", t, func() {

		Convey("função de geração de senhas retorna sempre mesmo valor para as mesmas entradas", func() {
			senha := "segredo"

			senhaGerada1 := GenerateSha1Password(senha)
			senhaGerada2 := GenerateSha1Password(senha)
			senhaGerada3 := GenerateSha1Password(senha)

			So(senhaGerada1, ShouldEqual, senhaGerada2)
			So(senhaGerada2, ShouldEqual, senhaGerada3)
		})

		Convey("resultado para uma senha vazia retorna vazio", func() {
			senhaGerada := GenerateSha1Password("")

			So(senhaGerada, ShouldBeEmpty)
		})
	})

	Convey("DOWNLOAD DE ARQUIVOS", t, func() {

		Convey("arquivo existente é recuperado com seus detalhes originais", func() {

			client := &http.Client{}
			req, err := http.NewRequest("GET", defaultVFolderHost+"files/7931cd97-6611-4e69-b3d4-dcbac4dba3c6", nil)
			configRequestAuth(req, "APP_1", defaultUserPassword)

			resp, err := client.Do(req)

			defer resp.Body.Close()

			So(resp, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, 200)
			So(resp.Header.Get(handlers.X_FILE_NAME_HEADER), ShouldEqual, "Jellyfish.jpg")
			So(resp.Header.Get(handlers.X_FILE_ID_HEADER), ShouldEqual, "7931cd97-6611-4e69-b3d4-dcbac4dba3c6")
		})

		Convey("arquivo não existe, nenhum resultado é retornado", func() {

			client := &http.Client{}

			//Esse arquivo não existe no repositório do usuário padrão
			req, err := http.NewRequest("GET", defaultVFolderHost+"files/7931cd97-6611-4e69-b3d4-dcbac4dba3c6", nil)
			configRequestAuth(req, defaultUserName, defaultUserPassword)

			resp, err := client.Do(req)

			defer resp.Body.Close()

			So(resp, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, 404)
		})
	})

	Convey("ESTATÍSTICAS", t, func() {

		Convey("recuperar estatisticas em json retorna resultados conhecidos", func() {

			client := &http.Client{}

			req, err := http.NewRequest("GET", defaultVFolderHost+"stats/stats.json", nil)
			configRequestAuth(req, "APP_2", defaultUserPassword)

			resp, err := client.Do(req)

			defer resp.Body.Close()

			So(resp.StatusCode, ShouldEqual, 200)
			So(resp, ShouldNotBeNil)
			So(err, ShouldBeNil)

			var stats model.VFStorageStats

			json.NewDecoder(resp.Body).Decode(&stats)

			So(stats.FileCount, ShouldEqual, 2)
		})

	})

	Convey("PAINEL DE CONTROLE", t, func() {

		Convey("página de painel de controle só é acessível ao admin", func() {
			client := &http.Client{}

			req1, err := http.NewRequest("GET", defaultHost+"/admin/ctrlpanel/", nil)
			configRequestAuth(req1, "APP_1", defaultUserPassword)

			resp1, err := client.Do(req1)

			defer resp1.Body.Close()

			So(resp1.StatusCode, ShouldEqual, 403)
			So(resp1, ShouldNotBeNil)
			So(err, ShouldBeNil)

			req2, err := http.NewRequest("GET", defaultHost+"/admin/ctrlpanel/", nil)
			configRequestAuth(req2, "admin", defaultUserPassword)

			resp2, err := client.Do(req2)

			defer resp2.Body.Close()

			So(resp2.StatusCode, ShouldEqual, 200)
			So(resp2, ShouldNotBeNil)
			So(err, ShouldBeNil)

		})
	})
}

//
// Funções utilitárias
//

func startServer() {
	config := ServerConfig{
		ServerPort:          defaultServerPort,
		DebugMode:           true,
		SharedFolder:        "./testdata/mystorage",
		TemplatesLocation:   "./templates",
		ServerUsersLocation: "./testdata/test.htpasswd"}
	Run(config)
}

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

func clearTempFiles() {
	os.RemoveAll("./testdata/mystorage/temp/")
}
