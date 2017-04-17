package players

type Player struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Level int    `json:"level"`
}

type Players []Player
