package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    net.JoinHostPort("localhost", "8080"),
		Handler: LoggingMiddleware(CreateServerMux()),
	}

	fmt.Printf("Server is listening on %s...\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
