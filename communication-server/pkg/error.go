package pkg

import "errors"

var (
	ErrEmailUnavailable = errors.New("system: email is unavailable")
)
