package etl

import (
	"encoding/json"
	"fmt"
	"github.com/EfosaE/naijacost-api/internal/util"
	"log"
	"strconv"
	"strings"
)

type CoHd struct {
	State string
	Cost  float64
}

// cohd = cost of a healthy diet
func LoadCoHdData() ([]byte, error) {

	sheet := util.Sheet{
		Filename:  "data/raw/CoHD_Nov_2024_Table.xlsx",
		SheetName: "cohd by national average",
	}
	rows, err := sheet.ReadSheet()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	s := util.Sheet{
		Rows: rows,
	}

	s.Print()
	var coHdData []CoHd
	// Skip header row
	for i, _ := range rows {
		row := rows[i]
		if len(row) < 2 {
			continue // Skip rows with insufficient data
		}

		state := strings.TrimSpace(row[0])
		if state == "" || strings.Contains(strings.ToLower(state), "national average") {
			continue // Skip empty rows or the national average row
		}

		// Clean up cost value: remove commas and convert to float
		costStr := strings.ReplaceAll(row[1], ",", "")
		cost, err := strconv.ParseFloat(costStr, 64)
		if err != nil {
			log.Printf("Warning: Could not parse cost for %s: %v", state, err)
			continue
		}

		coHdData = append(coHdData, CoHd{
			State: state,
			Cost:  cost,
		})
	}

	// Convert to JSON
	jsonData, err := json.Marshal(coHdData)
	if err != nil {
		log.Fatalf("Failed to convert to JSON: %v", err)
	}


	log.Printf("Processed %d entries for CoHd data", len(coHdData))

	return jsonData, nil

}
