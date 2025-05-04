package main

import (
	"log"
	"net/http"

	// "github.com/EfosaE/naijacost-api/internal/etl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)


type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}


func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/user", GetUser)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}
	render.JSON(w, r, user)
}