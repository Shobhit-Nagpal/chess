package handler

import (
	"fmt"
	"net/http"
)

func withHandlerPrefix(msg string) string {
	return "[handler]: " + msg
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func Analyze(w http.ResponseWriter, r *http.Request) {
	// TODO: write analyze logic here
	fmt.Fprintf(w, "Analyzed game")
}
