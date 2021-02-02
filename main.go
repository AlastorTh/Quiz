package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/AlastorTh/Quiz/models"
)

func main() {
	problems := readCSV("problems.csv")
	//fmt.Println(problems)
	var cont []models.Field

	for _, data := range problems {
		ans, err := strconv.Atoi(data[1])
		if err != nil {
			log.Fatalln(err)
		}
		cont = append(cont, models.Field{Question: string(data[0]), Answer: ans})
	}

	//fmt.Println(cont)

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
