package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"scibat/config"
	"scibat/router"
	"time"
)

func main() {
	config.Init()
	r := router.Init()
	serve(r)
}

func serve(r *mux.Router) {
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
