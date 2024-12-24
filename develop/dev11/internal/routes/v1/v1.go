package v1

import (
	"dev11/internal/routes/middlewares"
	"dev11/internal/view"
	"net/http"
)

type Router struct {
	View *view.View
}

func NewRouter(view *view.View) *Router {
	return &Router{
		View: view,
	}
}

func (r *Router) Events() http.Handler {
	v1mux := http.NewServeMux()

	v1mux.Handle("POST /create_event", middlewares.LogRequests(http.HandlerFunc(r.View.CreateEvent)))
	v1mux.Handle("POST /update_event", middlewares.LogRequests(http.HandlerFunc(r.View.UpdateEvent)))
	v1mux.Handle("POST /delete_event", middlewares.LogRequests(http.HandlerFunc(r.View.DeleteEvent)))

	v1mux.Handle("GET /events_for_day", middlewares.LogRequests(http.HandlerFunc(r.View.EventsForDay)))
	v1mux.Handle("GET /events_for_week", middlewares.LogRequests(http.HandlerFunc(r.View.EventsForWeek)))
	v1mux.Handle("GET /events_for_month", middlewares.LogRequests(http.HandlerFunc(r.View.EventsForMonth)))

	return http.StripPrefix("/v1", v1mux)
}
