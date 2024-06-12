package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReadCSV(file_path string) ([][]string, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func Process_data(records [][]string, column_name string) (float64, error) {
	if len(records) < 1 {
		return 0, fmt.Errorf("No records found")
	}

	header := records[0]
	col_index := -1
	for i, col := range header {
		if col == column_name {
			col_index = i
			break
		}
	}

	if col_index == -1 {
		return 0, fmt.Errorf("column %s	not found in header", column_name)
	}

	sum := 0.0
	for _, record := range records[1:] {
		value, err := strconv.ParseFloat(record[col_index], 64)
		if err != nil {
			return 0, err
		}
		sum += value

	}

	return sum, nil
}
