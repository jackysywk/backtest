package main

import (
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"time"
)
func Getdata(symbol string, start time.Time, end time.Time) (map[string]map[string]float64, error) {

	params := &chart.Params{
		Symbol:   symbol,
		Interval: datetime.OneDay,
		Start:     datetime.New(&start),
		End:       datetime.New(&end),
	}

	iter := chart.Get(params)
	data:= make(map[string]map[string]float64)
	for iter.Next() {
		bar := iter.Bar()
		bartime := time.Unix(int64(bar.Timestamp),0)
		formmatedTime:=bartime.Format("2006-01-02")
		if _, ok := data[formmatedTime]; !ok{
			data[formmatedTime] = make(map[string]float64)
		}
		data[formmatedTime]["Open"], _ = bar.Open.Float64()
		data[formmatedTime]["High"], _ = bar.High.Float64()
		data[formmatedTime]["Low"], _ = bar.Low.Float64()
		data[formmatedTime]["Close"], _ = bar.Close.Float64()
		data[formmatedTime]["Volume"] = float64(bar.Volume)
		
	}

	if iter.Err() != nil{
		return nil, iter.Err()
	}
	return data, nil
}
