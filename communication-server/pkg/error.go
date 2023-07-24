package pkg

import "errors"

var (
	ErrEmailUnavailable = errors.New("system: email is unavailable")
	ErrInvalidSignature = errors.New("system: invalid signature")
	ErrExpiredVersion   = errors.New("system: expired version")
)
