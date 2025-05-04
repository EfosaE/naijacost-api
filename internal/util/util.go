package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/aquasecurity/table"
	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	Filename  string
	SheetName string
	Rows      [][]string
}

func (s *Sheet) Print() {
	// Create a new table with headers
	t := table.New(os.Stdout)
	t.SetHeaders(s.Rows[0]...)
	for _, row := range s.Rows[1:] {
		t.AddRow(row...)
	}

	t.Render()
}

// ReadFromFile reads sheet rows from an already-opened file
func (s *Sheet) ReadFromFile(f *excelize.File) ([][]string, error) {
	rows, err := f.GetRows(s.SheetName)
	if err != nil {
		return nil, err
	}
	return cleanRows(rows), nil
}

// ReadSheet reads a new file 
func (s *Sheet) ReadSheet() ([][]string, error) {
	f, err := excelize.OpenFile(s.Filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(s.SheetName)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Clean the data (remove empty rows, trim whitespace)
	cleanedRows := cleanRows(rows)

	return cleanedRows, nil

}

// cleanRows removes empty rows and trims whitespace from each cell.
func cleanRows(rows [][]string) [][]string {
	var cleanedRows [][]string

	for _, row := range rows {
		if isEmptyRow(row) {
			continue
		}

		var cleanedRow []string
		for _, cell := range row {
			cleanedRow = append(cleanedRow, strings.TrimSpace(cell))
		}

		cleanedRows = append(cleanedRows, cleanedRow)
	}

	return cleanedRows
}

// isEmptyRow checks if a row contains only empty strings.
func isEmptyRow(row []string) bool {
	for _, cell := range row {
		if strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}
