package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	http.HandleFunc("/ping", ping)
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":80", nil)
}
