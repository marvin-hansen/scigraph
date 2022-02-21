package config_types

type DataChannelType uint8 // 255 possible values

const (
	UnknownDataChannelType DataChannelType = iota
	Book
	Trades
	Ticker
)

func (e DataChannelType) String() string {
	return [...]string{"UnknownDataChannelType", "Book", "Trades", "Ticker"}[e]
}
