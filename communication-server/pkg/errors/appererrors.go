package errors

import (
	"errors"
)

var (
	NotFound = errors.New("not_found")
	IlligalOperation = errors.New("illegal_operation")
	Internal = errors.New("internal")
)