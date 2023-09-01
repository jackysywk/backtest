import hkfdb
import sys




def get_hist_data(token, ticker, start=20220101, end=20230831, interval='1D'):
    client = hkfdb.Database(token)
    print("Downloading Data:",ticker)
    df = client.get_hk_stock_ohlc(ticker,start, end,interval,price_adj=True, vol_adj=True)
    df.to_csv(r'data/%s.csv'%ticker, encoding='utf-8')

if __name__ == "__main__":
    if len(sys.argv) <2:
        print("Usage: python get_data.py <ticker> [start date] [end date] [interval]")
        sys.exit(1)
    else:
        with open('token.txt','r') as file:
            token = file.read().strip()
        ticker = sys.argv[1]
        start_date = sys.argv[2] if len(sys.argv) > 2 else 20220101
        end_date = sys.argv[3] if len(sys.argv) > 3 else 20230831
        interval = sys.argv[4] if len(sys.argv) > 4 else '1D'
        
        get_hist_data(token, ticker, start_date, end_date, interval)
