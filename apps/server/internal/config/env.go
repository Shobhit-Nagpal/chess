package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	binaryPath []string
	port       string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Env{
		binaryPath: strings.Split(os.Getenv("BINARY_PATH"), ","),
		port:       fmt.Sprintf(":%s", os.Getenv("PORT")),
	}
}

func (e *Env) BinaryPath() []string {
	return e.binaryPath
}

func (e *Env) Port() string {
	return e.port
}
