package bookrp

import (
	"time"
)

type Book struct {
	Id       string
	Name     string
	Location string
	Size     int64
	Created  time.Time
}
