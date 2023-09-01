package main

import (
	"fmt"
	"sort"
	"time"
)

func backtest(data map[string]map[string]float64, candle_len float64, profit_target float64, stop_loss float64, holding_day float64) {
	open_date := time.Now()
	open_price, num_of_share, pnl, net_profit, num_of_trade := 0.0, 0.0, 0.0, 0.0, 0.0
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
		trade_logic := now_candle > candle_len

		//fmt.Printf("%v, %v\n", now_date_string, now_close)
		close_logic := now_date.Sub(open_date).Hours()/24 >= holding_day
		profit_target_cond := (now_close-open_price)/open_price > profit_target
		stop_loss_cond := (open_price-now_close)/open_price > stop_loss
		last_index_cond := i == len(data)
		//fmt.Printf("%v, Trade Logic:%v, Close Logic: %v, Profit: %v, Stop: %v\n", now_date, trade_logic, close_logic, profit_target_cond, stop_loss_cond)

		// Open Position
		if num_of_share == 0 && trade_logic {
			num_of_share = 1
			open_price = now_close
			open_date = now_date
			fmt.Printf("Open position: %v at %v\n", now_date_string, now_close)
		} else if num_of_share > 0 && (profit_target_cond || stop_loss_cond || last_index_cond || close_logic) {
			num_of_share = 0
			pnl = now_close - open_price
			net_profit += pnl
			num_of_trade += 1
			fmt.Printf("Close position: %v at %v\n", now_date_string, now_close)
		}

		i++
	}
	fmt.Printf("net_profit: %v\n", net_profit)
	fmt.Printf("num_of_trade: %v\n", num_of_trade)
}
