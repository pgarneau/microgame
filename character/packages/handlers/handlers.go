package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pgarneau/microgame/character/packages/character"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ListCharacters(w http.ResponseWriter, r *http.Request) {
	characters := character.Characters{
		character.Character{Name: "Aude", Class: "Mage", Level: 60},
		character.Character{Name: "Phil", Class: "Priest", Level: 60},
		character.Character{Name: "Oysto", Class: "Rogue", Level: 60},
	}

	if err := json.NewEncoder(w).Encode(characters); err != nil {
		panic(err)
	}
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	character := vars["character"]
	fmt.Fprintln(w, "Selected Player:", character)
}
