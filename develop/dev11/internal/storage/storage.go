package storage

import (
	"time"

	"go.uber.org/zap"
)

type Events struct {
	Log  *zap.Logger
	data []Event
}

type Event struct {
	UserId      int
	Date        time.Time
	Title       string
	Description string
}

func NewEvents(log *zap.Logger) *Events {
	return &Events{
		Log:  log,
		data: []Event{},
	}
}

func (s *Events) Save(event *Event) {
	s.data = append(s.data, *event)
}

func (s *Events) List() []Event {
	return s.data
}

func (s *Events) Replace(id int, event *Event) {
	s.data[id] = *event
}

func (s *Events) Delete(id int) {
	s.data = append(s.data[:id], s.data[id+1:]...)
}
