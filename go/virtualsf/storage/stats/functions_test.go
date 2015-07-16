package stats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContabilizarTamanhoDiretorioVazioRetornaZero(t *testing.T) {
	s, err := getStatsFromDirStorage("D:\\Temp\\virtualsf\\Vazia")

	assert.Nil(t, err, "Execução sem erros")
	assert.Equal(t, int64(0), s.TotalSize, "Tamanho computado deve ser ZERO")
	assert.Equal(t, 0, s.FileCount, "Quantidade de arquivos deve ser ZERO")
}

func TestContabilizarDirComArquivos(t *testing.T) {
	s, err := getStatsFromDirStorage("D:\\Temp\\virtualsf\\Temp")

	assert.Nil(t, err, "Execução sem erros")
	//assert.Equal(t, int64(0), s.TotalSize, "Tamanho computado deve ser ZERO")
	assert.Equal(t, 2, s.FileCount, "Quantidade de arquivos deve ser 2")
}
