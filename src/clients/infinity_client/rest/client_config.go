package rest

const HTTP_PROT = "http://"

type ClientConfig struct {
	Protocol         string
	Host             string
	Port             string
	connectionString string
}

func NewClientConfig(host, port string) *ClientConfig {
	return &ClientConfig{
		Host:             host,
		Port:             port,
		connectionString: HTTP_PROT + host + ":" + port + "/",
	}
}

func (c ClientConfig) GetConnectionString() string {
	return c.connectionString
}
