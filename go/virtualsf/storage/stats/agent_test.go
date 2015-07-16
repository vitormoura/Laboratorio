package stats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfiguracaoAgente(t *testing.T) {
	agent := NewAgent("D:\\Temp\\virtualsf", 1)

	assert.NotNil(t, agent, "Agente criado com sucesso")
	assert.True(t, true, "Dummy test")
}
