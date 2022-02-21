package config_types

type ExchangeName uint8

const (
	UnknownExchange ExchangeName = iota
	FTX
	VMX
)

func (e ExchangeName) String() string {
	return [...]string{"UNKNOWN_EXCHANGE", "FTX", "VMX"}[e]
}
