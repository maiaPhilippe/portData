package api

import (
	"clientAPI/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer(){

	mux := http.NewServeMux()

	fmt.Println("Start Server client ...")
	mux.HandleFunc("/id/", handlers.GetById)
	mux.HandleFunc("/upload", handlers.Upload)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-termChan

	log.Printf("shutting down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("down")
}
