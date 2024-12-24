package models

import domain "dev11/internal/domain/models"

func (m *CreateEvent) ToDomain() *domain.Event {
	return &domain.Event{
		UserId:      m.UserId,
		Date:        m.Date,
		Title:       m.Title,
		Description: m.Description,
	}
}

func (m *UpdateEvent) ToDomain() *domain.Event {
	return &domain.Event{
		EventId:     m.EventId,
		UserId:      m.UserId,
		Date:        m.Date,
		Title:       m.Title,
		Description: m.Description,
	}
}

func (m *DeleteEvent) ToDomain() *domain.Event {
	return &domain.Event{
		EventId: m.EventId,
	}
}

func (m *GetEvent) ToDomain() *domain.Event {
	return &domain.Event{
		UserId: m.UserId,
	}
}
