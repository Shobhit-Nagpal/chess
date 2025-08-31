package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		msg := fmt.Sprintf("[Method: %s] [Path: %s] [Duration: %s]\n", req.Method, req.URL.EscapedPath(), time.Since(start))
		slog.Info(msg)
	})
}
