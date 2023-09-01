package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	start := now.AddDate(-1, 0, 0)
	symbol := "00004"
	data, err := Getdata_hkfdb(symbol, start, now)
	//data, err := Getdata_yf(symbol, start, now)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}

	backtest(data, 1, 0.05, 0.2, 3)

}
