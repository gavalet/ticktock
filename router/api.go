package router

import (
	"cmd/ticktock/utils/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(address, port string) {
	log := logger.NewLogger("Initialise")
	router := LoadRoutes()
	server := &http.Server{
		Addr:    address + ":" + port,
		Handler: router,
	}
	go handleSystemSignals(server)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Info("Server error:", err)
	}
}

func handleSystemSignals(server *http.Server) {
	log := logger.NewLogger("Initialise")

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exit := make(chan int)
	go func() {

		s := <-signalChanel
		switch s {
		case syscall.SIGHUP:
			log.Info("SIGHUP received.")
			exit <- 0
		case syscall.SIGINT:
			log.Info("SIGINT received.")
			exit <- 0
		case syscall.SIGTERM:
			log.Info("SIGTERM received, put virtupian into maintenance mode before proceeding.")
			exit <- 0
		case syscall.SIGQUIT:
			log.Info("SIGQUIT received.")
			exit <- 0
		default:
			log.Info("Unknown signal received.")
			exit <- 1
		}
	}()
	exitCode := <-exit
	shutdownServer(server)
	os.Exit(exitCode)
}

func shutdownServer(server *http.Server) {
	log := logger.NewLogger("Initialise")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Stop accepting new connections.
	server.SetKeepAlivesEnabled(false)
	err := server.Shutdown(ctx)
	if err != nil {
		log.Error("Error during server shutdown: %v", err)
	}
}
