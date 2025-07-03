package main

import (
	"encoding/csv"
	"os"
)

func Import_from_csv (csv_file string) ([][]string, int, error) {

	file, err := os.Open(csv_file)
	if err != nil {
		return nil, 0, err
	}

	defer file.Close()

	var reader *csv.Reader = csv.NewReader(file)

	file_arr, err := reader.ReadAll()
	if err != nil {
		return file_arr, len(file_arr), err
	}

	return file_arr, len(file_arr), nil
}
