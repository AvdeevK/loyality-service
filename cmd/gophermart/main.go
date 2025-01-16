package main

import (
	logger "github.com/AvdeevK/loyality-service/internal/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

func run(r chi.Router) error {
	logger.Log.Info("service successfully started at port 8080")
	return http.ListenAndServe(":8080", r)
}

func main() {
	if err := logger.Init("info"); err != nil {
		panic(err)
	}
	srv := chi.NewRouter()
	if err := run(srv); err != nil {
		logger.Log.Fatal("service isn't strted, fatal error")
		os.Exit(1)
	}
}
