package models

import "dev11/internal/storage"

func (e *Event) ToStorage() *storage.Event {
	return &storage.Event{
		UserId:      e.UserId,
		Date:        e.Date,
		Title:       e.Title,
		Description: e.Description,
	}
}

func ToDomain(e *storage.Event) *Event {
	return &Event{
		UserId:      e.UserId,
		Date:        e.Date,
		Title:       e.Title,
		Description: e.Description,
	}
}
