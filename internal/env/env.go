package env

// Base URL: https://api.weather.gov for NWS API

import (
	"fmt"
)

type Env struct {
	Port       string
	BaseUrlNWS string
}

func InitEnv() *Env {
	fmt.Println("initEnv invoked")
	e := &Env{
		Port:       "8080",
		BaseUrlNWS: "https://api.weather.gov",
	}
	return e
}

// would usually have some type of function here to load value's from OS or vault or kubernetes configMap using helm
// would have port := os.Getenv("PORT") in there for example
