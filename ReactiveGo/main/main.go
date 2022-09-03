package main

import (
	"ReactiveGo/reactive"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Reactive WebSocket")
	reactive.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
