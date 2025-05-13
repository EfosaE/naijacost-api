package etl

import (
	"context"
	"fmt"
	"strconv"

	"github.com/EfosaE/naijacost-api/internal/db"
	"github.com/EfosaE/naijacost-api/internal/db/sqlc"
	"github.com/EfosaE/naijacost-api/internal/util"
)

type StatesService struct {
	db *db.DB
}

func NewStatesService(db *db.DB) *StatesService {
	return &StatesService{
		db: db,
	}
}

func ReadStatesCostDataFromFile() ([][]string, error) {
	sheet := util.Sheet{
		Filename:  "data/raw/transport_costs_by_state.xlsx",
		SheetName: "state transport",
	}

	rows, err := sheet.ReadSheet()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sheet.Rows = rows
	// sheet.Print()
	return rows, nil
}



func (s *StatesService) SetStatesCostDataIntoDB(ctx context.Context) (int64, error) {
	// Read the states cost data from the file
	rows, err := ReadStatesCostDataFromFile()
	if err != nil {
		fmt.Println("Error reading states cost data:", err)
		return 0, err
	}

	// Create a slice to hold all the parameter structs
	params := make([]sqlc.BulkInsertStateCostsParams, 0, len(rows))

	// Skip the header row if it exists
	startIdx := 0
	if len(rows) > 0 && rows[0][0] == "State" {
		startIdx = 1
	}

	// Process each data row
	for i := startIdx; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 6 {
			fmt.Printf("Skipping row %d: insufficient columns\n", i)
			continue
		}

		// Get state name
		stateName := row[0]

		// Convert string values to float64
		airCost, err := strconv.ParseFloat(row[1], 8)
		if err != nil {
			fmt.Printf("Error parsing Air_Cost for %s: %v\n", stateName, err)
			continue
		}

		busCostInter, err := strconv.ParseFloat(row[2], 8)
		if err != nil {
			fmt.Printf("Error parsing Bus_Cost_Inter for %s: %v\n", stateName, err)
			continue
		}

		busCostIntra, err := strconv.ParseFloat(row[3], 8)
		if err != nil {
			fmt.Printf("Error parsing Bus_Cost_Intra for %s: %v\n", stateName, err)
			continue
		}

		motorcycleCost, err := strconv.ParseFloat(row[4], 8)
		if err != nil {
			fmt.Printf("Error parsing Motorcycle_Cost for %s: %v\n", stateName, err)
			continue
		}

		waterCost, err := strconv.ParseFloat(row[5], 8)
		if err != nil {
			fmt.Printf("Error parsing Water_Cost for %s: %v\n", stateName, err)
			continue
		}

		// Create parameter struct for this row
		param := sqlc.BulkInsertStateCostsParams{
			State:          stateName,
			AirCost:        airCost,
			BusCostInter:   busCostInter,
			BusCostIntra:   busCostIntra,
			MotorcycleCost: motorcycleCost,
			WaterCost:      waterCost,
		}

		// Add to our slice of parameters
		params = append(params, param)
	}



	// fmt.Println(params)
	result, err := s.db.Queries.BulkInsertStateCosts(ctx, params)

	if err != nil {
		fmt.Println("Error inserting states cost data:", err)
		return 0, err
	}

	fmt.Println("Successfully inserted states cost data:", result)
	return result, nil
}


func (s *StatesService) GetStatesCostData(ctx context.Context) ([]sqlc.StatesCost, error) {
	// Fetch the state costs from the database
	data, err := s.db.Queries.ListStatesCosts(ctx)
	if err != nil {
		fmt.Println("Error fetching states cost data:", err)
		return nil, err
	}

	return data, nil
}