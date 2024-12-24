package domain

import (
	"dev11/internal/storage"

	"go.uber.org/zap"
)

type Domain struct {
	Log     *zap.Logger
	Storage *storage.Events
}

func NewDomain(log *zap.Logger) *Domain {
	return &Domain{
		Log:     log,
		Storage: storage.NewEvents(log),
	}
}
