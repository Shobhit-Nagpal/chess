package routes

import (
	"net/http"

	"github.com/Shobhit-Nagpal/chess/apps/server/internal/server/handler"
	"github.com/Shobhit-Nagpal/chess/apps/server/internal/server/middleware"
)

func RegisterHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.Index)
	mux.HandleFunc("GET /health", handler.HealthCheck)

	mux.HandleFunc("POST /analyze", handler.Analyze)

	handler := middleware.Logger(mux)

	return handler
}
