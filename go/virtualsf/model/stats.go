package model

import (
	"time"
)

//VFStorageStats apresenta informações sobre estatísticas de armazenamento de arquivos
type VFStorageStats struct {
	App       string
	Date      time.Time
	TotalSize int64
	FileCount int
}
