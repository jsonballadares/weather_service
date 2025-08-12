package env

import "net/http"

// Base URL: https://api.weather.gov for NWS API

type Env struct {
	Port       string
	BaseUrlNWS string
	HttpClient *http.Client
}

func InitEnv() *Env {
	e := &Env{
		Port:       "8080",
		BaseUrlNWS: "https://api.weather.gov",
		HttpClient: http.DefaultClient,
	}
	return e
}

// would usually have some type of function here to load value's from OS or vault or kubernetes configMap using helm
// would have port := os.Getenv("PORT") in there for example
