package etl

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/EfosaE/naijacost-api/internal/db/sqlc"
	"github.com/EfosaE/naijacost-api/internal/util"
)

// cohd = cost of a healthy diet
func readCoHdData() ([]sqlc.BulkInsertStateFoodCostsParams, error) {

	sheet := util.Sheet{
		Filename:  "data/raw/CoHD_Nov_2024_Table.xlsx",
		SheetName: "cohd by national average",
	}
	rows, err := sheet.ReadSheet()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sheet.Rows = rows

	sheet.Print()
	var coHdData []sqlc.BulkInsertStateFoodCostsParams
	// Skip header row
	for i := range rows {
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

		coHdData = append(coHdData, sqlc.BulkInsertStateFoodCostsParams{
			State: state,
			Cost:  util.ToFloat8(cost),
		})
	}

	return coHdData, nil
}

func (s *StatesService) SetCoHdDataIntoDB(ctx context.Context) (int64, error) {
	// Read the states cost data from the file
	coHdData, err := readCoHdData()
	if err != nil {
		fmt.Println("Error reading states cost data:", err)
		return 0, err
	}

	result, err := s.db.Queries.BulkInsertStateFoodCosts(ctx, coHdData)
	if err != nil {
		fmt.Println("Error inserting states food costs data:", err)
		return 0, err
	}

	fmt.Println("Successfully inserted states food costs data:", result)
	return result, nil
}
