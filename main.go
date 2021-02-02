package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	problems := readCSV("problems.csv")
	fmt.Println(problems)
}

func readCSV(filepath string) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	reader.Comment = '#'

	record, err := reader.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	return record
}
