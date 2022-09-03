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
	_, cancel := start()
	defer shutdown(cancel)
	WaitShutdown()
}

func start() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(context.Background())
	Start()
	return
}

func WaitShutdown() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-sigc
	log.Printf("signal received [%v] canceling everything\n", s)
}

func shutdown(cancel context.CancelFunc) {
	cancel()
	ctx, cancelTimeOut := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelTimeOut()
	doneHTTP := Shutdown(ctx)
	err := waitUntilIsDoneOrCanceled(ctx, doneHTTP)
	if err != nil {
		log.Printf("service stopped by timeout %s\n", err)
	}
	log.Println("Shutdown completed")
}

func waitUntilIsDoneOrCanceled(ctx context.Context, dones ...chan struct{}) (err error) {
	done := make(chan struct{})
	go func() {
		for _, d := range dones {
			<-d
		}
		close(done)
	}()

	select {
	case <-done:
		log.Println("All done")
	case <-ctx.Done():
		err = http.ErrServerClosed
		log.Println("canceled")
	}
	return
}
