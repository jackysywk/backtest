package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// FileExists checks if a file exists and is not a directory.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Getdata_hkfdb(ticker string, start time.Time, end time.Time) (map[string]map[string]float64, error) {
	start_string := start.Format("20060102")
	end_string := end.Format("20060102")
	filename := fmt.Sprintf("data/%s.csv", ticker)
	if !FileExists(filename) {
		fmt.Println("Downloading ticker data using python")
		cmd := exec.Command("python", "get_data.py", ticker, start_string, end_string)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to execute command: %s", err)
		}
	} else {
		fmt.Println("Using cache file.")
	}
	data := make(map[string]map[string]float64)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		outerKey := record[0]

		// Ensure an inner map has been created for the outer key
		if _, ok := data[outerKey]; !ok {
			data[outerKey] = make(map[string]float64)
		}

		// Iterate over the remaining columns
		for i := 1; i < len(header); i++ {
			value, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				log.Fatal(err)
			}

			data[outerKey][header[i]] = value
		}
	}

	return data, err
}

//func main() {Getdata_hkfdb("AAPL")}
