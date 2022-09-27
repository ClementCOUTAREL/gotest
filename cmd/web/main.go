package main

import (
	"fmt"
	"net/http"

	"github.com/ClementCOUTAREL/pkg/handlers"
)

const portNumber = "127.0.0.1:8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/contact", handlers.Contact)

	fmt.Println(fmt.Sprintf("Starting application on port %v", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
