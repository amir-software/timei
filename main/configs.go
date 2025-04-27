package main

import (
	"net/http"
	"time"
)

type config struct {
	addr string
}

type application struct {
	config config
}

// func (app *application) mount() http.Handler {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)

// 	r.Route("/v1", func(r chi.Router) {
// 		r.Get("/health", app.healthCheckHandler)
// 	})

// 	return r
// }

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	return srv.ListenAndServe()
}
