package main

import (
	"fmt"

	"github.com/jasonballadares/weather_service/api/routers"
	"github.com/jasonballadares/weather_service/internal/env"
)

func main() {
	e := env.InitEnv()

	r := routers.InitRouter()

	err := routers.Serve(r, e.Port)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error starting server: %s", err))
	}
}
