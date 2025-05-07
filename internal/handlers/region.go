package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/EfosaE/naijacost-api/internal/apierror"
	"github.com/go-chi/render"
)

// GeopoliticalZone represents a geopolitical zone in Nigeria with its states
type GeopoliticalZone struct {
	Name   string   `json:"name"`
	States []string `json:"states"`
}

// NigeriaData represents the full structure of Nigeria's geopolitical data
type NigeriaData struct {
	GeopoliticalZones []GeopoliticalZone `json:"geopolitical_zones"`
}

// GetStates handles requests for Nigeria's states
func GetStates(w http.ResponseWriter, r *http.Request) {
	// Read the JSON file
	fileData, err := os.ReadFile("data/raw/nigeria_zones_states.json")
	if err != nil {
		fmt.Println(err)
		render.Render(w, r, apierror.InternalServerError(err, "Failed to read states data"))
		return
	}


	// Parse the JSON data
	var nigeriaData NigeriaData
	if err := json.Unmarshal(fileData, &nigeriaData); err != nil {
		fmt.Println(err)
		// render.Status(r, http.StatusInternalServerError)
		// render.JSON(w, r, map[string]string{
		// 	"error":   "Failed to parse states data",
		// 	"message": err.Error(),
		// })
		render.Render(w, r, apierror.InternalServerError(err, "Failed to parse states data"))
		return
	}

	// // Extract all states from all zones
	// var allStates []string
	// for _, zone := range nigeriaData.GeopoliticalZones {
	// 	allStates = append(allStates, zone.States...)
	// }

	// // Sort the states alphabetically
	// sort.Strings(allStates)

	// Return the states as JSON
	render.JSON(w, r, nigeriaData)
}