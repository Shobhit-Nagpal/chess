package utils

import (
	"log"
	"os"
	"path"
)

var binaryPath = []string{"personal", "oss", "Stockfish", "src", "stockfish"}

func GetStockfishPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("[error]: could not get home dir")
		return ""
	}

	fullPath := append([]string{homeDir}, binaryPath...)

	return path.Join(fullPath...)
}
