package character

type Character struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Level int    `json:"level"`
}

type Characters []Character
