package main

import (
	"github.com/pgarneau/microgame/player/packages/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
