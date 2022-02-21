package config_types

// Timeframe specification: https://docs.ftx.com/#get-historical-prices
type Timeframe string

const (
	TIME_15_sec   Timeframe = "15"
	TIME_01_min   Timeframe = "60"
	TIME_03_min   Timeframe = "300"
	TIME_15_min   Timeframe = "900"
	TIME_01_hour  Timeframe = "3600"
	TIME_04_hour  Timeframe = "14400"
	TIME_01_Day   Timeframe = "86400"
	TIME_03_Day   Timeframe = "259200"
	TIME_01_Week  Timeframe = "604800"
	TIME_02_Week  Timeframe = "1209600"
	TIME_01_Month Timeframe = "2592000" // 30*86400 = 30 DAYS(!)
)

func (t Timeframe) String() string {
	return string(t)
}
