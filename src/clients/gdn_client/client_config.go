package gdn_client

type ClientConfig struct {
	Protocol         string
	Host             string
	Port             string
	Timeout          int
	connectionString string
}

func NewClientDefaultConfig() *ClientConfig {
	host := ""
	port := ""
	return &ClientConfig{
		Host:             host,
		Port:             port,
		Timeout:          5,
		connectionString: HTTPS_PROT + host + ":" + port,
	}
}

func NewClientConfig(host, port string, timeout int) *ClientConfig {
	return &ClientConfig{
		Host:             host,
		Port:             port,
		Timeout:          timeout,
		connectionString: HTTPS_PROT + host + ":" + port,
	}
}

func (c ClientConfig) GetConnectionString() string {
	return c.connectionString
}
