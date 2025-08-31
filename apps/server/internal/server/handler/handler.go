package handler

import (
	"fmt"
	"log/slog"
	"net/http"
)

func withHandlerPrefix(msg string) string {
	return "[handler]: " + msg
}

func Index(w http.ResponseWriter, r *http.Request) {
	slog.Info(withHandlerPrefix("/ hit!"))
	fmt.Fprintf(w, "Welcome to my website!")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	slog.Info(withHandlerPrefix("/health hit!"))
	fmt.Fprintf(w, "OK")
}

func Analyze(w http.ResponseWriter, r *http.Request) {
	slog.Info(withHandlerPrefix("/analyze hit!"))
	// TODO: write analyze logic here
	fmt.Fprintf(w, "Analyzed game")
}
