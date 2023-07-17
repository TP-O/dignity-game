package errors

import (
	"errors"
)

var (
	InvalidInput = errors.New("invalid_input")
	EmptyInput = errors.New("empty_input")
)