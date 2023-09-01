# backtest

## The 2nd commit

Add data source of HKFDB.  excluding input data parameter to hkfdb.
HKFDB will use the file in data folder, which will make date parameter effectiveless.  If you want to redownload the csv, you may need to delete all csv file in data folder.

### future plan
- Multi-thread backtesting in golang
- flexible strategy writer
- output the backtest file.