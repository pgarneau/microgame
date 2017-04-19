package model

type Character struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Level int    `json:"level"`
	ID    int    `json:"id"`
}

type Characters []Character
