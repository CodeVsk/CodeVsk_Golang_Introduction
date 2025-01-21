package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 5)

		return
	})

	server := &http.Server{Addr: ":8080"}

	go func() {
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatal(err)
		}
	}()

		GracefulllShuwtdown(server)
}

func GracefulllShuwtdown(server *http.Server){
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-shutdown 

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}