package routers

// basic skeleton of a router package
// using chi as I am familiar with it and use it a good bit at work
// could've just used net/http or gorilla/mux or fiber
// I like it as it feels like net/http++ and has lots of neat features like url param support
// middleware support and integrates with context well which I didn't use but would useful at scale

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasonballadares/weather_service/api/handlers"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/weather", handlers.WeatherHandler)
	return r
}

func Serve(r *chi.Mux, port string) error {
	fmt.Println(fmt.Sprintf("Starting server on port: %s", port))
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		return err
	}
	return nil
}
