package model

import (
	"time"
)

//VFStorageStats apresenta informações sobre estatísticas de armazenamento de arquivos
type VFStorageStats struct {
	App       string
	Date      time.Time
	TotalSize float32
	FileCount int
}
