package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUnknown        = errors.New("unknown error")
	ErrLowBalance     = errors.New("low balance")
	ErrRowExist       = errors.New("row exist")
	ErrRowNotExist    = errors.New("row doesn't exist")
)
