package main

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"nuri-challenge/controllers"
	"time"
)

func configureHTTPDefaultTransport() {
	// configure the default http transport to allow many idle connections per host
	// by default only two are allowed (see http.DefaultMaxIdleConnsPerHost)
	if tp, ok := http.DefaultTransport.(*http.Transport); ok {
		tp.MaxIdleConnsPerHost = 100
		tp.MaxIdleConns = 500
	}
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Info("Panic in go-git initialization.", "error", err)
		}
	}()
	configureHTTPDefaultTransport()
	e := controllers.InitEcho()
	controllers.AddRoutes(e)
	s := http.Server{
		Addr:        ":8080",
		Handler:     e,
		ReadTimeout: 30 * time.Second,
	}
	log.Info("Server Started on 8080")
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
