package domain

import "errors"

var (
	ErrAlreadyExist = errors.New("Already exists.")
	ErrNotExist     = errors.New("Not exists.")
	ErrInvalidInput = errors.New("invalid input")
)
