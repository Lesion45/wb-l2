package middlewares

import (
	"dev11/pkg/logger"
	"net/http"

	"go.uber.org/zap"
)

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.NewReqZap()

		log.Info("incoming request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)

		next.ServeHTTP(w, r)
	})
}
