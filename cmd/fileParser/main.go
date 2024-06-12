package main

import (
	"fmt"
	"log"

	"github.com/MandavkarPranjal/fileParser/internal/parser"
)

func main() {
	filepath := "data.csv"
	column_name := "amount"

	records, err := parser.ReadCSV(filepath)
	if err != nil {
		log.Fatalf("Error reading csv file: %v", err)
	}

	sum, err := parser.Process_data(records, column_name)
	if err != nil {
		log.Fatalf("Error processing data: %v", err)
	}

	fmt.Printf("Sum of column '%s': %f\n", column_name, sum)
}
