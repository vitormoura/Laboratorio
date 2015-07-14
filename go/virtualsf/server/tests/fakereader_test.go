package tests

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFakeFileReaderInterface(t *testing.T) {

	fake := NewFakeReader(1024)
	reader := bufio.NewReader(fake)
	content, _ := reader.ReadString(' ')

	assert.Equal(t, 1024, len([]byte(content)), "Conteudo da string precisa ter 32000 caracteres")
}
