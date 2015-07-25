package logs

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type logInfo struct {
	appID   string
	logName string
	msg     string
}

//ServerLog é um componente capaz de registrar eventos do servidor
type ServerLog struct {
	location string
	file     io.WriteCloser
	logger   *log.Logger

	writerC chan *logInfo
}

func (s *ServerLog) listen() {
	go func() {
		for l := range s.writerC {
			s.logger.Println("\t", l.appID, "\t", l.logName, "\t", l.msg)
		}
	}()
}

//Dispose cancela execução do logger
func (s *ServerLog) Dispose() {
	close(s.writerC)
	s.file.Close()
}

//Info registra uma nova entrada do tipo informativo
func (s *ServerLog) Info(appID string, logName string, msg string) {
	s.writerC <- &logInfo{appID, logName, msg}
}

//New cria uma nova instância do log capaz de registrar operações do servidor
func New(dir string) (*ServerLog, error) {

	location := filepath.Join(dir, "vfolder.log")
	file, err := os.OpenFile(location, os.O_WRONLY|os.O_CREATE, os.ModeAppend)

	if err != nil {
		return nil, err
	}

	log := log.New(file, "", log.LstdFlags)
	srv := &ServerLog{location, file, log, make(chan *logInfo)}
	srv.listen()

	return srv, nil

}
