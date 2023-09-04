package main

import (
	"fmt"
	"sort"
	"time"
)

func backtest(data map[string]map[string]float64, account Account, params Params) {
	// Initialize the capital
	account.capital = account.initial_capital

	account.lot_size = 100
	open_date := time.Now()
	// i is declared to iterate a sequential map
	i := 0
	keys := make([]time.Time, 0, len(data))
	for k := range data {
		t, _ := time.Parse("2006-01-02", k)
		keys = append(keys, t)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})
	for _, now_date := range keys {
		now_date_string := now_date.Format("2006-01-02")
		now_close := data[now_date_string]["Close"]
		now_open := data[now_date_string]["Open"]
		now_candle := now_close - now_open

		// Calculate the equity
		account.unrealized_pnl = account.num_of_share * (now_close - account.open_price)
		account.equity_value = account.capital + account.unrealized_pnl
		account.net_profit = account.equity_value - account.initial_capital

		trade_logic := now_candle > params.candle_len

		//fmt.Printf("%v, %v\n", now_date_string, now_close)
		close_logic := now_date.Sub(open_date).Hours()/24 >= params.holding_day
		profit_target_cond := (now_close-account.open_price)/account.open_price > params.profit_target
		stop_loss_cond := (account.open_price-now_close)/account.open_price > params.stop_loss
		last_index_cond := i == len(data)
		//fmt.Printf("%v, Trade Logic:%v, Close Logic: %v, Profit: %v, Stop: %v\n", now_date, trade_logic, close_logic, profit_target_cond, stop_loss_cond)

		// Open Position
		if account.num_of_share == 0 && trade_logic {
			account.num_of_share = get_buyable_share(account.capital, now_close, account.lot_size)
			account.open_price = now_close
			open_date = now_date
			fmt.Printf("Open position: %v at %v\n", now_date_string, now_close)
			// Close Position
		} else if account.num_of_share > 0 && (profit_target_cond || stop_loss_cond || last_index_cond || close_logic) {

			account.realized_pnl = account.unrealized_pnl
			account.num_of_share = 0
			account.capital += account.realized_pnl
			account.num_of_trade += 1
			fmt.Printf("Close position: %v at %v\n", now_date_string, now_close)
		}

		i++
	}
	fmt.Printf("net_profit: %v\n", account.net_profit)
	fmt.Printf("num_of_trade: %v\n", account.num_of_trade)
}
