package utils

import (
	"log"
	"os"
	"path"

	"github.com/Shobhit-Nagpal/chess/apps/server/internal/config"
)

func GetStockfishPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("[error]: could not get home dir")
		return ""
	}

	var binaryPath = config.GetConfig().GetEnv().BinaryPath()

	fullPath := append([]string{homeDir}, binaryPath...)

	return path.Join(fullPath...)
}
