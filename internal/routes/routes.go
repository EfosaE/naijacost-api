package routes

import (
	"github.com/EfosaE/naijacost-api/internal/db"
	"github.com/EfosaE/naijacost-api/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func StatesRouter(r chi.Router, db *db.DB) {
    r.Post("/", handlers.SetStatesCostDataHandler(db))
	r.Get("/list", handlers.GetStatesCostDataHandler(db))
}

// func CoHdRouter(r chi.Router) {
// 	r.Get("/list", handlers.GetCoHdList)
// }
