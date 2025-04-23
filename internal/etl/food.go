package etl

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func LoadFoodPrices() {
	// This function will be called by the ETL process to load food prices into the system
	// The implementation will depend on the specific requirements and data sources used in the ETL process
	// It involves reading data from an Xlsx file, transforming it, and loading it into a database
	f, err := excelize.OpenFile("Book1.xlsx")
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

	fmt.Println(f)

}
