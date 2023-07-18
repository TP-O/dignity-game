package helpers

import "github.com/google/uuid"

type UIDGen interface {
	New() string
}

type uidgen struct {}

func NewUIDGen() UIDGen {
	return &uidgen{}
}

func (u *uidgen) New() string {
	return uuid.New().String()
}