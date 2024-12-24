package routes

import (
	v1 "dev11/internal/routes/v1"
	"dev11/internal/view"
	"net/http"
)

func InitRouter(view *view.View) *http.ServeMux {
	mux := http.NewServeMux()
	v1 := v1.NewRouter(view)

	mux.Handle("/v1/", v1.Events())

	mux.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	return mux
}
