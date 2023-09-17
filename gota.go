package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/markcheno/go-talib"
	"os"
)

func main() {
	csvfile, err := os.Open("data/00001.csv")
	if err != nil {
		fmt.Println(err)
	}
	df := dataframe.ReadCSV(csvfile)

	close := df.Col("Close").Float()
	rsi := talib.Rsi(close, 9)
	rsiSeries := series.Floats(rsi)
	rsiSeries.Name = "RSI"
	df = df.Mutate(rsiSeries)
	//fmt.Println(df.Select([]string{"datetime", "Open", "High", "low", "Close", "RSI"}))
	// Convert DataFrame to Records
	records := df.Records()

	// Iterate over the records and print each one
	fmt.Println(records[0])
	for _, record := range records[9:19] {
		fmt.Println(record[3], record[4], record[5], record[8])
	}
	/*
		n := 20
		indices := make([]int, n)
		for i := 0; i < n; i++ {
			indices[i] = i
		}
		fmt.Println(df.Subset(indices))
	*/
	//fmt.Println(df.Select([]string{"Close"}))
}
