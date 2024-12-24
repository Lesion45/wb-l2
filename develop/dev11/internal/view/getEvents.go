package view

import (
	"dev11/internal/view/models"
	"net/http"
	"time"
)

func (v *View) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	userId := r.FormValue("user_id")

	dto, err := models.NewGetEvent(userId)
	if err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	events := models.NewEventResponse(v.Domain.GetEventInRange(dto.ToDomain(), time.Hour*24))

	if err := events.Send(&w, http.StatusOK); err != nil {
		v.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (v *View) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	userId := r.FormValue("user_id")

	dto, err := models.NewGetEvent(userId)
	if err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	events := models.NewEventResponse(v.Domain.GetEventInRange(dto.ToDomain(), time.Hour*24*7))

	if err := events.Send(&w, http.StatusOK); err != nil {
		v.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (v *View) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	userId := r.FormValue("user_id")

	dto, err := models.NewGetEvent(userId)
	if err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	events := models.NewEventResponse(v.Domain.GetEventInRange(dto.ToDomain(), time.Hour*24*7*30))

	if err := events.Send(&w, http.StatusOK); err != nil {
		v.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
