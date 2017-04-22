package controller

import (
	"app/model"
	"app/shared/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Waddup, %q", html.EscapeString(r.URL.Path))
}

func ListCharacters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	model.ListDB()

	if err := json.NewEncoder(w).Encode(database.AllCharacters); err != nil {
		panic(err)
	}

}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var character model.Character
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &character); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	c := database.CreateCharacter(character)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	characterName := vars["name"]
	character := database.GetCharacter(characterName)
	fmt.Fprintln(w, "Selected Character:", character.Name)
	fmt.Fprintln(w, "Character Class:", character.Class)
	fmt.Fprintln(w, "Character Level:", character.Level)
}
