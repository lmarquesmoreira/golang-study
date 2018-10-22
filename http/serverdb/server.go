package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", UsuarioHandler)
	log.Println("running")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
