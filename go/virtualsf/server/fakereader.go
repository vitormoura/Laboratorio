package server

import (
	"io"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//fakeReader é uma implementação de io.Reader para permitir a
//simulação de envio de arquivos com bytes gerados aleatoriamente
type fakeReader struct {
	size    int
	currPos int
}

func (f *fakeReader) getBytesRemaining() int {
	return f.size - f.currPos
}

func (f *fakeReader) Read(p []byte) (n int, err error) {

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

//NewFakeReader cria um novo io.Reader capaz de recuperar uma
//quantidade definida de bytes aleatórios
func NewFakeReader(totalBytesToRead int) io.Reader {
	return &fakeReader{size: totalBytesToRead, currPos: 0}
}
