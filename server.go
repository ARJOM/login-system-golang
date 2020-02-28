package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login/", LoginHandler)
	http.HandleFunc("/register/", RegisterHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
