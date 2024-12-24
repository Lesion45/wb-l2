package models

import (
	"dev11/internal/domain/models"
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Error: err.Error(),
	}
}

func (e *ErrorResponse) Send(w *http.ResponseWriter, status int) error {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	return json.NewEncoder(*w).Encode(e)
}

type EventResponse struct {
	Result []models.Event `json:"result"`
}

func NewEventResponse(events []models.Event) *EventResponse {
	return &EventResponse{
		Result: events,
	}
}

func (e *EventResponse) Send(w *http.ResponseWriter, status int) error {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	return json.NewEncoder(*w).Encode(e)
}
