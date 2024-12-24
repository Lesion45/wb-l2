package main

import (
	"dev11/internal/routes"
	"dev11/internal/view"
	"net/http"

	"dev11/pkg/logger"
)

func main() {
	log := logger.NewAppZap()
	view := view.NewView(log)

	mux := routes.InitRouter(view)

	log.Info("server started on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err.Error())
	}
}
