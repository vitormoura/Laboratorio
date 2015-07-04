package model

import (
	"fmt"
	"io"
	"time"
)

//FileInfo oferece informações básicas sobre um arquivo já armazenado
type FileInfo struct {
	ID   string
	Name string
}

//File representa um arquivo que pode ser armazenado
type File struct {
	ID          string
	App         string
	Name        string
	Size        int64
	PublishDate time.Time
	MimeType    string
	Properties  map[string]string
	Stream      io.ReadCloser `json:"-"`
}

func (f File) String() string {

	return fmt.Sprintf(
		`--------------------------------
Name: %s
Size: %d
App: %s
PublishDate: %v
Properties: %v
--------------------------------`,
		f.Name, f.Size, f.App, f.PublishDate, f.Properties)
}
