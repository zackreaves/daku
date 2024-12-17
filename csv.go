package main

import (
	"encoding/csv"
	"os"
)

func Import_from_csv (csv_file string) [][]string {
	file, err := os.Open(csv_file)
	defer file.Close()
	Error_check(err)
	var reader *csv.Reader = csv.NewReader(file)
	file_arr, err := reader.ReadAll()
	Error_check(err)

	return file_arr
}
