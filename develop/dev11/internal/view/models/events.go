package models

import (
	"strconv"
	"time"
)

type CreateEvent struct {
	UserId      int
	Date        time.Time
	Title       string
	Description string
}

func NewCreateEvent(userId string, date string, title string, description string) (*CreateEvent, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, NewErrInvalidData(err.Error())
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, NewErrInvalidData(err.Error())
	}

	if title == "" || description == "" {
		return nil, NewErrInvalidData("title or description is empty")
	}

	return &CreateEvent{
		UserId:      userIdInt,
		Date:        parsedDate,
		Title:       title,
		Description: description,
	}, nil
}

type UpdateEvent struct {
	EventId     int
	UserId      int
	Date        time.Time
	Title       string
	Description string
}

func NewUpdateEvent(eventId string, userId string, date string, title string, description string) (*UpdateEvent, error) {
	var parsedDate time.Time
	var userIdInt int = -1
	var err error

	if eventId == "" {
		return nil, NewErrInvalidData("event_id is empty")
	}
	eventIdInt, err := strconv.Atoi(eventId)
	if err != nil {
		return nil, NewErrInvalidData(err.Error())
	}

	if date != "" {
		parsedDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return nil, NewErrInvalidData(err.Error())
		}
	}

	if userId != "" {
		userIdInt, err = strconv.Atoi(userId)
		if err != nil {
			return nil, NewErrInvalidData(err.Error())
		}
	}

	return &UpdateEvent{
		EventId:     eventIdInt,
		UserId:      userIdInt,
		Date:        parsedDate,
		Title:       title,
		Description: description,
	}, nil
}

type DeleteEvent struct {
	EventId int
}

func NewDeleteEvent(eventId string) (*DeleteEvent, error) {
	if eventId == "" {
		return nil, NewErrInvalidData("event_id is empty")
	}
	eventIdInt, err := strconv.Atoi(eventId)
	if err != nil {
		return nil, NewErrInvalidData(err.Error())
	}
	return &DeleteEvent{
		EventId: eventIdInt,
	}, nil
}

type GetEvent struct {
	UserId int
}

func NewGetEvent(userId string) (*GetEvent, error) {
	if userId == "" {
		return nil, NewErrInvalidData("user_id is empty")
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, NewErrInvalidData(err.Error())
	}
	return &GetEvent{
		UserId: userIdInt,
	}, nil
}
