package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	binaryPath []string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Env{
		binaryPath: strings.Split(os.Getenv("BINARY_PATH"), ","),
	}
}

func (e *Env) BinaryPath() []string {
	return e.binaryPath
}
