package etl

import (
	"fmt"
	"os"

	"github.com/aquasecurity/table"
	"github.com/xuri/excelize/v2"
)

func LoadFoodPrices() {
	// Create a new table with headers
	t := table.New(os.Stdout)

	// This function will be called by the ETL process to load food prices into the system
	// The implementation will depend on the specific requirements and data sources used in the ETL process
	// It involves reading data from an Xlsx file, transforming it, and loading it into a database
	f, err := excelize.OpenFile("data/raw/selected_food_oct_2024.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("selected food oct 2024")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.SetHeaders(rows[0]...)
	for _, row := range rows[1:] {
		t.AddRow(row...)
	}

	t.Render()
}
