package etl

import (
	"fmt"
	"github.com/EfosaE/naijacost-api/internal/util"
)

// type StatesCost struct {
// 	State string
// 	Cost  float64
// }

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


