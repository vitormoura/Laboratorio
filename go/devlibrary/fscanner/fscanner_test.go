package fscanner

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	_ "github.com/vitormoura/Laboratorio/go/devlibrary/bookrp"
	"testing"
)

//
// Configuração
//

type FScannerTestSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(FScannerTestSuite))
}

func (s *FScannerTestSuite) SetupTest() {
}

//
//Testes
//

func (s *FScannerTestSuite) TestReconheceCorretamenteArquivosDiretorioValido() {

	qtdeLivrosEncontrados := 0

	for b := range ScanBooksFromDir("./test_data/docs") {
		fmt.Println(b.Name)
		qtdeLivrosEncontrados++
	}

	assert.Equal(s.T(), 4, qtdeLivrosEncontrados)
}
