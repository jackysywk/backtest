package main

import (
	"math"
)

func clip(value, min, max float64) float64 {
	return math.Min(max, math.Max(min, value))
}
func round(x float64, prec int) float64 {
	pow := math.Pow(10, float64(prec))
	return math.Round(x*pow) / pow
}

func commission(price float64) float64 {
	//平台費
	var platform_fee float64 = 15

	//佣金
	commission := clip(price*0.003/100, 3, 100)
	//印花稅
	tax := price * 0.13 / 100
	if tax < 1 {
		tax = 1
	} else {
		tax = math.Ceil(tax)
	}
	//交易費
	cost := price * 0.00565 / 100
	if cost < 0.01 {
		cost = 0.01
	}
	//證監會交易徵
	other_fee := round(price*0.000027, 2)
	if other_fee < 0.01 {
		other_fee = 0.01
	}
	//財務滙報局交易微費
	extra := price * 0.00015 / 100

	total_commission := round(platform_fee+commission+tax+other_fee+cost+extra, 2) * 2
	return total_commission

}
