package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EfosaE/naijacost-api/internal/etl"
	"github.com/EfosaE/naijacost-api/internal/apierror"
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
	port := os.Getenv("PORT")
	fmt.Println("Port:", port)
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	etl.LoadCoHdData()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome To My Cost Analyser!"))
	})

	r.NotFound(apierror.NotFoundHandler())
	r.MethodNotAllowed(apierror.MethodNotAllowedHandler())

	// Creating a Sub Router
	api := chi.NewRouter()
	api.Route("/states", routes.StatesRouter)
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
