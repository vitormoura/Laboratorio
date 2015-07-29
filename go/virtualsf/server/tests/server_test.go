package tests

import (
	"bytes"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server"
	"github.com/vitormoura/Laboratorio/go/virtualsf/server/handlers"
	_ "log"
	"net/http"
	"os"
	"testing"
)

const (
	defaultHost         = "http://localhost:4045"
	defaultVFolderHost  = defaultHost + "/vfolder/"
	defaultUserName     = "temp"
	defaultServerPort   = 4045
	defaultUserPassword = "segredo"
)

func startServer() {
	config := server.ServerConfig{
		ServerPort:          defaultServerPort,
		DebugMode:           true,
		SharedFolder:        "./testdata/mystorage",
		TemplatesLocation:   "../templates",
		ServerUsersLocation: "./testdata/test.htpasswd"}
	server.Run(config)
}

func clearTempFiles() {
	os.RemoveAll("./testdata/mystorage/temp/")
}

func TestServer(t *testing.T) {

	startServer()

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

			req1, err := http.NewRequest("GET", defaultHost+"/ctrlpanel/", nil)
			configRequestAuth(req1, "APP_1", defaultUserPassword)

			resp1, err := client.Do(req1)

			defer resp1.Body.Close()

			So(resp1.StatusCode, ShouldEqual, 404)
			So(resp1, ShouldNotBeNil)
			So(err, ShouldBeNil)

			req2, err := http.NewRequest("GET", defaultHost+"/ctrlpanel/", nil)
			configRequestAuth(req2, "admin", defaultUserPassword)

			resp2, err := client.Do(req2)

			defer resp2.Body.Close()

			So(resp2.StatusCode, ShouldEqual, 200)
			So(resp2, ShouldNotBeNil)
			So(err, ShouldBeNil)

		})
	})
}
