package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfiguracaoAgente(t *testing.T) {
	agent := NewStatsUpdateAgent(testStorageRootDir, 1)

	assert.NotNil(t, agent, "Agente criado com sucesso")
}
