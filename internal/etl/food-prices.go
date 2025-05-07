package etl

import (
	"fmt"

	"time"

	"github.com/EfosaE/naijacost-api/internal/util"
)

func LoadFoodPrices() {
	// This function will be called by the ETL process to load food prices into the system
	// The implementation will depend on the specific requirements and data sources used in the ETL process
	// It involves reading data from an Xlsx file, transforming it, and loading it into a database

	start := time.Now()

	loadAvgPriceState()
	loadAvgPriceRegion()

	fmt.Println("Took:", time.Since(start))
}

func loadAvgPriceState() {

	sheet := util.Sheet{
		Filename:  "data/raw/selected_food_Dec_2024.xlsx",
		SheetName: "selected food Dec 2024",
	}
	rows, err := sheet.ReadSheet()
	if err != nil {
		fmt.Println(err)
		return
	}

	s := util.Sheet{
		Rows: rows,
	}
	s.Print()
}

func loadAvgPriceRegion() {

	sheet := util.Sheet{
		Filename:  "data/raw/selected_food_Dec_2024.xlsx",
		SheetName: "zone all item",
	}
	rows, err := sheet.ReadSheet()
	if err != nil {
		fmt.Println(err)
		return
	}

	s := util.Sheet{
		Rows: rows,
	}
	s.Print()
}
