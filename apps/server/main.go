package main

import (
	"log"

	"github.com/Shobhit-Nagpal/chess/apps/server/internal/engine"
	"github.com/Shobhit-Nagpal/chess/apps/server/internal/uci"
)

func main() {
	engine := engine.New()

	if err := engine.Spawn(); err != nil {
		log.Fatal(err)
	}

	go func() {
		for engine.IsRunning() {
			response := engine.ReadResponse()
			uci.ParseUciCommand(response)
		}
	}()

	engine.SendCommand(uci.IsReady)
	engine.SendCommand(uci.Uci)

	for {
	}
}
