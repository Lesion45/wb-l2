package models

import "time"

type Event struct {
	EventId     int
	UserId      int
	Date        time.Time
	Title       string
	Description string
}
