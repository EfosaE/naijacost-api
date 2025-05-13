package handlers

import (
	"net/http"

	"github.com/EfosaE/naijacost-api/internal/api"
	"github.com/EfosaE/naijacost-api/internal/db"
	"github.com/EfosaE/naijacost-api/internal/etl"
	"github.com/go-chi/render"
)

// SetStatesCostDataHandler handles uploading state cost data
func SetStatesCostDataHandler(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Create the service with the DB
		statesService := etl.NewStatesService(db)
		result, err := statesService.SetStatesCostDataIntoDB(ctx)

		if err != nil || result == 0 {
			render.Render(w, r, api.InternalServerError(err, "Failed to set states cost data"))

			return
		}

		api.SendSuccess(w, r, api.OK(result, "States cost data set successfully"))
	}

}


// GetStatesCostDataHandler handles getting state cost data
func GetStatesCostDataHandler(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Create the service with the DB
		statesService := etl.NewStatesService(db)
		result, err := statesService.GetStatesCostData(ctx)

		if err != nil {
			render.Render(w, r, api.InternalServerError(err, "Failed to get states cost data"))

			return
		}

		api.SendSuccess(w, r, api.OK(result, "States cost data retrieved successfully"))
	}
}