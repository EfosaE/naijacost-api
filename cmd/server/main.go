package main

import (
	"context"
	"log"
	"net/http"

	"github.com/EfosaE/naijacost-api/internal/api"
	"github.com/EfosaE/naijacost-api/internal/config"
	"github.com/EfosaE/naijacost-api/internal/db"

	// "github.com/EfosaE/naijacost-api/internal/etl"
	"github.com/EfosaE/naijacost-api/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {

	ctx := context.Background()
	config.Load()
	db, err := db.InitDB(ctx)

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Pool.Close()

	port := config.App.Port
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome To My Cost Analyser!"))
	})

	r.NotFound(api.NotFoundHandler())
	r.MethodNotAllowed(api.MethodNotAllowedHandler())

	// Creating a Sub Router
	api := chi.NewRouter()
	api.Route("/states", func(r chi.Router) {
		routes.StatesRouter(r, db)
	})
	api.Route("/cohd", routes.CoHdRouter)

	// apparently this is also a valid way to create a sub router
	// api/v1/states
	// routes.StatesRouter(api)

	// Mounting the new Sub Router on the main router
	r.Mount("/api/v1", api)

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// func GetStates(w http.ResponseWriter, r *http.Request) {
// 	states := []string{"Lagos", "Abuja", "Port Harcourt"}
// 	render.JSON(w, r, states)
// }
