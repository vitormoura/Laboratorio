package bookrp

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type books struct {
	dsn string
	db  *sql.DB
}

func (b *books) isOpen() bool {
	return b.db != nil
}

//Open abre conectividade com a origem de dados
func (b *books) Open() error {
	var err error

	b.db, err = sql.Open("sqlite3", b.dsn)

	return err
}

//Close encerra conectividade com origem de dados
func (b *books) Close() {
	b.db.Close()
}

//Add adiciona um novo livro ao repositório
func (b *books) Add(book Book) bool {

	if !b.isOpen() {
		log.Panic("Repositorio não aberto")
	}

	_, err := b.db.Exec("INSERT INTO CAD_BOOKS ( COD_BOOK, NOM_BOOK, TXT_PATH_BOOK, DTH_CADASTRO, QTD_BYTES_BOOK ) VALUES ( ?, ?, ?, ?, ? )", book.Id, book.Name, book.Location, book.Created, book.Size)

	return err == nil
}

//Count recupera quantidade de livros presentes no repositório
func (b *books) Count() int {
	var qtde int

	err := b.db.QueryRow("SELECT COUNT(*) FROM CAD_BOOKS").Scan(&qtde)

	if err != nil {
		log.Panic("Erro ao recuperar quantidade de livros")
	}

	return qtde
}

//Remove exclui o livro identificado pelo ID informado
func (b *books) Remove(id string) bool {

	if !b.isOpen() {
		log.Panic("Repositorio não aberto")
	}

	_, err := b.db.Exec("DELETE FROM CAD_BOOKS WHERE COD_BOOK = ?", id)

	return err == nil
}

//New cria uma nova instância do repositório de livros
func New(dbFilePath string) *books {

	b := new(books)
	b.dsn = dbFilePath

	return b
}
