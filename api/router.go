package main

import "net/http"

func CreateServerMux() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{ "message": "Hello from server!" }`))
	})

	return router
}
