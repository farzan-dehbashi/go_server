package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tilt-dev/tilt-example-go/internal/handler"
)

func main() {
	http.Handle("/", handler.NewExampleRouter())
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}
