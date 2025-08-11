package env

type Env struct {
	Port string
}

func initEnv() *Env {
	e := &Env{
		Port: "8080",
	}
	return e
}

// would usually have some type of function here to load value's from OS or vault or kubernetes configMap using helm
