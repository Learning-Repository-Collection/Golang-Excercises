package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// opening the text file
	file, err := os.Open("problems.csv")

	// checking for errors
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// closing the file
	defer file.Close()

	// reading from the file
	reader := csv.NewReader(file)

	// reads all records from the .csv file
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	for _, eachRecord := range records {
		fmt.Println(eachRecord)
	}








}