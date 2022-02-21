package config_types

type ChannelType uint8 // 255 possible values

const (
	UnknownChannelType ChannelType = iota
	AmxAccountChannel
	ApxApiChannel
	CmxControlChannel
	ClientDmxDataChannel
	EmxOrderExecutionChannel
)

func (e ChannelType) String() string {
	return [...]string{"UnknownChannelType", "AmxAccountChannel", "ApxApiChannel", "CmxControlChannel", "ClientDmxDataChannel", "EmxOrderExecutionChannel"}[e]
}
