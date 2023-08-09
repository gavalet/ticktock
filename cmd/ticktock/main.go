package main

import (
	"cmd/ticktock/router"
	l "cmd/ticktock/utils/logger"
	"os"
)

func main() {
	log := l.NewLogger("Initialise")

	if len(os.Args) < 2 {
		log.Error("Usage: go run main.go <listen_addr>:<port>")
		os.Exit(1)
	}
	addr := os.Args[1]
	port := os.Args[2]
	log.Debug("port: ", port)
	log.Debug("address: ", addr)

	log.Printf("Lets start")
	router.Run(addr, port)
}
