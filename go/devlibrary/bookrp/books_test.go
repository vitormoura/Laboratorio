package bookrp

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

//
// Configuração
//

const (
	DBPATH         = "./sample.db"
	DBCREATESCRIPT = "./db.sql"
)

type BooksTestSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(BooksTestSuite))
}

func (s *BooksTestSuite) SetupTest() {

	os.Remove(DBPATH)

	bytes, _ := ioutil.ReadFile(DBCREATESCRIPT)
	sqlStm := string(bytes)

	db, _ := sql.Open("sqlite3", DBPATH)
	db.Exec(sqlStm)

	defer db.Close()
}

//
//Testes
//

func (s *BooksTestSuite) TestInclusaoNovoLivroValidoRetornaSucesso() {
	repo := New(DBPATH)
	defer repo.Close()

	b1 := Book{Name: "Um livro exemplo", Location: "c:\\Livros\\b1.pdf", Created: time.Now(), Id: "123456", Size: 10}

	repo.Open()
	repo.Add(b1)

	assert.Equal(s.T(), 1, repo.Count())
}

func (s *BooksTestSuite) TestExclusaoLivroValidoRetornaSucesso() {
	repo := New(DBPATH)
	defer repo.Close()

	b1 := Book{Name: "Um livro exemplo", Location: "c:\\Livros\\b1.pdf", Created: time.Now(), Id: "123456", Size: 10}

	repo.Open()
	repo.Add(b1)
	repo.Remove("123456")

	assert.Equal(s.T(), 0, repo.Count())
}
