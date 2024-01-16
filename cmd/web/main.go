package main

import (
	"fmt"
	"net/http"

	"github.com/rtmelsov/mansProducts/pkg/handlers"
)

// port number to work on develop
const portNumber = ":8080"

// main is the main application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprint("Starting the application on port", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
