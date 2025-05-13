package etl

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestReadStatesCostDataFromFile(t *testing.T) {
	// Change working dir to project root
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(b))) // 3 levels up from internal/etl
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}

	rows, err :=  ReadStatesCostDataFromFile()

	// Check for error
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if rows == nil {
		t.Fatal("Expected rows to be non-nil, got nil")
	}

	if len(rows) == 0 {
		t.Fatal("Expected rows to contain data, got empty slice")
	}

	// Optional: check structure of data
	for i, row := range rows {
		if i == 0 {
			expectedHeaders := []string{"State", "Air Cost", "Bus Cost Inter", "Bus Cost Intra", "Motorcycle Cost", "Water Cost"} // Adjust as needed
			if !reflect.DeepEqual(row, expectedHeaders) {
				t.Errorf("Expected headers %v, got %v", expectedHeaders, row)
			}
		}
	}
}
