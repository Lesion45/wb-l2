package view

import (
	"dev11/internal/view/models"
	"net/http"
)

func (v *View) CreateEvent(w http.ResponseWriter, r *http.Request) {
	date := r.FormValue("date")
	userId := r.FormValue("user_id")
	title := r.FormValue("title")
	description := r.FormValue("description")

	event, err := models.NewCreateEvent(userId, date, title, description)
	if err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if err := v.Domain.SaveEvent(event.ToDomain()); err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusServiceUnavailable); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
}

func (v *View) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	eventId := r.FormValue("event_id")
	date := r.FormValue("date")
	userId := r.FormValue("user_id")
	title := r.FormValue("title")
	description := r.FormValue("description")

	event, err := models.NewUpdateEvent(eventId, date, userId, title, description)
	if err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if err := v.Domain.UpdateEvent(event.ToDomain()); err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusServiceUnavailable); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (v *View) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	eventId := r.FormValue("event_id")

	event, err := models.NewDeleteEvent(eventId)
	if err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusBadRequest); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if err := v.Domain.DeleteEvent(event.ToDomain()); err != nil {
		v.Log.Error(err.Error())
		if err := models.NewErrorResponse(err).Send(&w, http.StatusServiceUnavailable); err != nil {
			v.Log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
