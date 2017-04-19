package main

import (
	"app/route"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
