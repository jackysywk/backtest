package main

import (
	"math"
)

func get_buyable_share(capital, stock_price, lot_size float64) float64 {
	num_of_lot := math.Floor(capital / (lot_size * stock_price))
	buyable_share := num_of_lot * lot_size
	return buyable_share
}

/*
func main() {
	buyable_share := get_buyable_share(10500, 10, 100)
	fmt.Println(buyable_share)

}
*/
