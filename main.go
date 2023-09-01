package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	start := now.AddDate(-2, 0, 0)
	symbol := "AAPL"
	data, err := Getdata_yf(symbol, start, now)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
	backtest(data, 3, 0.05, 0.2, 3)

}
