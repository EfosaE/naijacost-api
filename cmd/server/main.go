package main

import (
	"log"
	"net/http"
	"os"

	// "github.com/EfosaE/naijacost-api/internal/etl"
	"github.com/EfosaE/naijacost-api/internal/apierror"
	"github.com/EfosaE/naijacost-api/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// etl.LoadFoodPrices()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.NotFound(apierror.NotFoundHandler())
	r.MethodNotAllowed(apierror.MethodNotAllowedHandler())

	// Creating a Sub Router
	apiRouter := chi.NewRouter()
	apiRouter.Get("/states", handlers.GetStates)

	// Mounting the new Sub Router on the main router
	r.Mount("/api/v1", apiRouter)

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// func GetStates(w http.ResponseWriter, r *http.Request) {
// 	states := []string{"Lagos", "Abuja", "Port Harcourt"}
// 	render.JSON(w, r, states)
// }
