# backtest

#### The 2nd commit

Add data source of HKFDB.  excluding input data parameter to hkfdb.
HKFDB will use the file in data folder, which will make date parameter effectiveless.  If you want to redownload the csv, you may need to delete all csv file in data folder.

#### future plan
- Multi-thread backtesting in golang
- flexible strategy writer
- output the backtest file.

#### 20230904
- add the commission func
- but if i want to implement SMA / other RSA technical analysis. the current price data structure is not supporting.
- I need to rebuild the price data structure into a slice containing struct.