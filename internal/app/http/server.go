package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/LIYINGZHEN/go-consul-example/internal/app/service"
	"github.com/gorilla/mux"
)

type AppServer struct {
	Service *service.Service
}

// Run will start the http server.
func (a *AppServer) Run(port string) {
	r := mux.NewRouter()
	a.publicRoutes(r)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("%v", err)
		}
	}()

	log.Printf("app is running at localhost:%s", port)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	srv.Shutdown(context.Background())
	log.Printf("shutting down app")
	os.Exit(0)
}
