package config_types

type ServiceName uint8

const (
	UNKNOWN ServiceName = iota
	QUESTDB
	TSDB
	NATS
	AMDB
	CMDB
	SMDB
	IMDB
	TMX
	IMX_FTX
)

func (d ServiceName) String() string {
	return [...]string{"UNKNOWN", "questDB", "tsdb", "nats", "amdb", "cmdb", "smdb", "imdb", "tmx", "imx_ftx"}[d]
}

type ServiceExposureType uint8

const (
	ENDPOINT ServiceExposureType = iota // 0 :
	CHANNEL
)

func (d ServiceExposureType) String() string {
	return [...]string{"Endpoint", "Channel"}[d]
}
