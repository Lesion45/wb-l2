package view

import (
	"dev11/internal/domain"

	"go.uber.org/zap"
)

type View struct {
	Domain *domain.Domain
	Log    *zap.Logger
}

func NewView(log *zap.Logger) *View {
	return &View{
		Domain: domain.NewDomain(log),
		Log:    log,
	}
}
