package main

import (
	"context"
	"log"
	"net/http"
)

var srv *http.Server

func Start() {
	createServer()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func Shutdown(ctx context.Context) (done chan struct{}) {
	done = make(chan struct{})
	go func() {
		defer close(done)
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("couldnt shutdown the server: [%s]\n", err)
		}
	}()
	return
}

func createServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	srv = &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello World!"))
	if err != nil {
		log.Printf("coulnt write response: [%s]\n", err)
	}
}
