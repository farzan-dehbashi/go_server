package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tilt-dev/tilt-example-go/internal/handler"
)

func setupHandlers() {
	http.Handle("/", handler.NewWebRouter())
}
func main() {
	setupHandlers()
	err := http.ListenAndServe("", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}
