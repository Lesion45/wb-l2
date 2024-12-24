package domain

import (
	"dev11/internal/domain/models"
	"dev11/internal/storage"
	"dev11/pkg/collections"
	"time"
)

func (d *Domain) SaveEvent(event *models.Event) error {
	d.Storage.Save(event.ToStorage())
	return nil
}

func (d *Domain) UpdateEvent(event *models.Event) error {
	list := d.Storage.List()

	if event.EventId < 0 || event.EventId >= len(list) {
		d.Log.Error("invalid event id")
		return models.NewErrInvalidData("invalid event id")
	}

	if event.Date.IsZero() {
		event.Date = list[event.EventId].Date
	}
	if event.UserId == -1 {
		event.UserId = list[event.EventId].UserId
	}
	if event.Title == "" {
		event.Title = list[event.EventId].Title
	}
	if event.Description == "" {
		event.Description = list[event.EventId].Description
	}

	d.Storage.Replace(event.EventId, event.ToStorage())

	return nil
}

func (d *Domain) DeleteEvent(event *models.Event) error {
	d.Storage.Delete(event.EventId)
	return nil
}

func (d *Domain) GetEventInRange(event *models.Event, r time.Duration) []models.Event {
	list := d.Storage.List()

	events := collections.Filter(list, func(e storage.Event) bool {
		return e.Date.Sub(time.Now()) <= r && e.UserId == event.UserId
	})

	var result []models.Event
	for idx, e := range events {
		ev := models.ToDomain(&e)
		ev.EventId = idx
		result = append(result, *ev)
	}

	return result
}
