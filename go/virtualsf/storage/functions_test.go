package storage

import (
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vitormoura/Laboratorio/go/virtualsf/model"
	"testing"
)

const testStorageRootDir = "D:\\Temp\\virtualsf-tests\\"

func TestContabilizarDirComArquivos(t *testing.T) {

	resultC, _ := getStatsFromDirStorage(testStorageRootDir)

	stats := make([]model.VFStorageStats, 0, 10)

	for stat := range resultC {
		stats = append(stats, stat)
	}

	assert.Equal(t, 2, len(stats), "Quantidade de arquivos deve ser 2")
	assert.Equal(t, "SYS-A", stats[0].App, "Primeiro sistema e SYS-A")
	assert.Equal(t, "SYS-V", stats[1].App, "Segundo sistema e SYS-V")

	assert.Equal(t, 2, stats[0].FileCount, "SYS-A deve possuir dois arquivos")
	assert.Equal(t, 0, stats[1].FileCount, "SYS-V nao deve possuir arquivos")

}
