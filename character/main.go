package main

import (
	"github.com/pgarneau/microgame/character/packages/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
