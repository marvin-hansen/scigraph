package rest

func getEndpoint(config *ClientConfig) string {
	if config == nil {
		return DefaultEndpoint
	} else {
		return config.GetConnectionString()
	}
}
