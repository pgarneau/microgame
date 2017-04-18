package database

import (
	"fmt"
	"github.com/pgarneau/microgame/character/packages/character"
)

var currentID int

var AllCharacters character.Characters

// Give us some OG data
func init() {
}

func GetCharacter(name string) character.Character {
	for _, c := range AllCharacters {
		if c.Name == name {
			return c
		}
	}
	// return empty Character if not found
	return character.Character{}
}

func CreateCharacter(c character.Character) character.Character {
	currentID += 1
	c.ID = currentID
	AllCharacters = append(AllCharacters, c)
	return c
}

func DeleteCharacter(id int) error {
	for i, c := range AllCharacters {
		if c.ID == id {
			AllCharacters = append(AllCharacters[:i], AllCharacters[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Character with id of %d to delete", id)
}
