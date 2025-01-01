package main

// all required imports
import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)


// function to open the csv file
func readCSVFile(filename string) ([] byte, error) {

	// opening the file 
	fileContents, openErrorMessage := os.Open(filename)

	// checking if an error is thrown
	if openErrorMessage != nil {
		return nil, openErrorMessage
	}

	// closing the file
	defer fileContents.Close()

	// reading all the contents of the file
	fileData, readErrorMessage := io.ReadAll(fileContents)

	// checking if an error is thrown
	if readErrorMessage != nil {
		return nil, readErrorMessage
	}

	// returning the content and nil
	return fileData, nil

}


// function to parse through the csv
// converts raw CSV data into a csv.Reader object
func parseCSV(data []byte) (*csv.Reader, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	return reader, nil
}


// function to process the csv.Reader object
func processCSV(reader *csv.Reader) {

	// initiating a for loop
	for {

		// reading a single row from the .csv file
		record, errorMessage := reader.Read()

		// checking if the end of the file is reached
		if errorMessage == io.EOF {
			break

		// checking if the file has not been yet read	
		} else if errorMessage != nil {
			fmt.Println("Error Reading CSV Data: ", errorMessage)
			break
		}


		// iterating through each letter in the string
		

		// printing out the contents of the row
		fmt.Println(record)

	}
}


// the main function
func main() {

	// reading contents of the CSV file
	fileData, errorMessage := readCSVFile("problems.csv")

	// if there is an error in terms of reading the .csv file+
	if errorMessage != nil {
		fmt.Println("Error Reading File: ", errorMessage)
		return
	}

	// parsing the CSV
	reader, errorMessage := parseCSV(fileData)

	if errorMessage != nil {
		fmt.Println("Error creating CSV reader: ", errorMessage)
		return
	}

	// calling the processCSV method
	processCSV(reader)
	

}

