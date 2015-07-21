package model

import (
	"time"
)

//VFStorageStats apresenta informações sobre estatísticas de armazenamento de arquivos
type VFStorageStats struct {
	Date      time.Time
	TotalSize int64
	FileCount int
}
