package main

import (
	"fmt"
	"time"
)

type Account struct {
	initial_capital float64
	capital         float64
	equity_value    float64
	open_price      float64
	num_of_share    float64
	unrealized_pnl  float64
	realized_pnl    float64
	num_of_trade    float64
	net_profit      float64
	lot_size        float64
}

type Params struct {
	candle_len    float64
	candle_dir    string
	profit_target float64
	stop_loss     float64
	holding_day   float64
}

func main() {

	// Get Data from API
	now := time.Now()
	start := now.AddDate(-4, 0, 0)
	symbol := "00388"
	data, err := Getdata_hkfdb(symbol, start, now)
	//data, err := Getdata_yf(symbol, start, now)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
	//Initialize the Account Profile
	account := Account{
		initial_capital: 100000,
		equity_value:    0,
		capital:         0,
		open_price:      0,
		num_of_share:    0,
		unrealized_pnl:  0,
		realized_pnl:    0,
		num_of_trade:    0,
		net_profit:      0,
		lot_size:        0,
	}
	//Initialize the Parameter
	params := Params{
		candle_len:    5,
		candle_dir:    "negative",
		profit_target: 0.05,
		stop_loss:     0.05,
		holding_day:   10,
	}

	backtest(data, account, params)

}
