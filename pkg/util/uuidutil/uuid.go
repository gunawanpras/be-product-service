package uuidutil

import (
	"github.com/google/uuid"
)

type UUID interface {
	New() uuid.UUID
}

type uuidHelper struct {
}

func (h *uuidHelper) New() uuid.UUID {
	return uuid.New()
}

var UUIDHelper UUID = &uuidHelper{}
