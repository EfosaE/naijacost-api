package handlers

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/EfosaE/naijacost-api/internal/api"
// 	"github.com/EfosaE/naijacost-api/internal/etl"
// 	"github.com/go-chi/render"
// )

// Cohd (cost of healthy diet) represents the cost of a healthy diet in Nigeria
type CoHd struct {
	State string  `json:"state"`
	Cost  float64 `json:"cost"`
}

// // GetCoHdList handles requests for the cost of a healthy diet in Nigeria
// func GetCoHdList(w http.ResponseWriter, r *http.Request) {
// 	// Read the JSON file
// 	jsonData, err := etl.LoadCoHdData()
// 	if err != nil {
// 		log.Println(err)
// 		render.Render(w, r, api.InternalServerError(err, "Failed to read states data"))
// 		return
// 	}

// 	// Parse the JSON data
// 	var coHdData []CoHd
// 	if err := json.Unmarshal(jsonData, &coHdData); err != nil {
// 		log.Println(err)
// 		render.Render(w, r, api.InternalServerError(err, "Failed to parse states data"))
// 		return
// 	}

// 	// Return the coHdData as JSON
// 	render.JSON(w, r, coHdData)
// }


