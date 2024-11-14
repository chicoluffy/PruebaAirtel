package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading file")
	}
	validRecords := [][]string{}
	for _, record := range records {
		if !isEmptyRecord(record) {
			validRecords = append(validRecords, record)
		}
	}
	return validRecords, nil
}

func isEmptyRecord(record []string) bool {
	for _, value := range record {
		if value == "" {
			return true
		}
	}
	return false
}
