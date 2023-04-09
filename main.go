package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/prestonchoate/go-user-service/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.GetRoot)
	mux.HandleFunc("/hello", handlers.GetHello)
	mux.HandleFunc("/login", handlers.HandleLogin)
	mux.HandleFunc("/users", handlers.GetUsers)

	err := http.ListenAndServe(":3333", mux)

	if (errors.Is(err, http.ErrServerClosed)) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server %s\n", err)
		os.Exit(1)
	}
}