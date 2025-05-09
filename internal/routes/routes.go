package routes

import (
	"github.com/EfosaE/naijacost-api/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func StatesRouter(r chi.Router) {
	r.Get("/", handlers.GetStates)
}

func CoHdRouter(r chi.Router) {
	r.Get("/list", handlers.GetCoHdList)
}
