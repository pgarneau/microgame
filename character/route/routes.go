package route

import (
	"github.com/pgarneau/microgame/character/controller"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var AllRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"ListCharacters",
		"GET",
		"/characters",
		controller.ListCharacters,
	},
	Route{
		"CreateCharacter",
		"POST",
		"/characters",
		controller.CreateCharacter,
	},
	Route{
		"Character",
		"GET",
		"/characters/{name}",
		controller.GetCharacter,
	},
}
