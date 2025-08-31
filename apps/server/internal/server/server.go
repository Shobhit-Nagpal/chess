package server

import (
	"log/slog"
	"net/http"

	"github.com/Shobhit-Nagpal/chess/apps/server/internal/config"
	"github.com/Shobhit-Nagpal/chess/apps/server/internal/server/routes"
)

func ListenAndServe() {
	port := config.GetConfig().GetEnv().Port()

	handler := routes.RegisterHandler()

	server := &http.Server{
		Handler: handler,
		Addr:    port,
	}

	slog.Info("Server is up and running!")
	server.ListenAndServe()
}
