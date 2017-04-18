package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pgarneau/microgame/player/packages/players"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ListPlayers(w http.ResponseWriter, r *http.Request) {
	players := players.Players{
		players.Player{Name: "Aude", Class: "Mage", Level: 60},
		players.Player{Name: "Phil", Class: "Priest", Level: 60},
		players.Player{Name: "Oysto", Class: "Rogue", Level: 60},
	}

	if err := json.NewEncoder(w).Encode(players); err != nil {
		panic(err)
	}
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["player"]
	fmt.Fprintln(w, "Selected Player:", player)
}
