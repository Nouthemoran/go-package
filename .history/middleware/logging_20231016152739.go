package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"URL": r.Host + r.URL.Path,
			"METHOD":   10,
		}).Info("A group of walrus emerges from the ocean")
		next.ServeHTTP(w, r)
	})
}
