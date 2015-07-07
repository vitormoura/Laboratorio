package model

import (
	"errors"
)

var (
	ErrFileNotFound         = errors.New("virtualfs: File not found")
	ErrFileNotSupported     = errors.New("virtualfs: File type not supported")
	ErrStorageFail          = errors.New("virtualfs: Storage error")
	ErrStorageInvalidConfig = errors.New("virtualfs: Storage invalid configuration")
	ErrEmptyFile            = errors.New("virtualfs: File is empty")
	ErrStorageLocked        = errors.New("virtualfs: Storage locked")
)
