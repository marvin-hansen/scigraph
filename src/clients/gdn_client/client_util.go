package gdn_client

import "time"

func getTimeout(config *ClientConfig) time.Duration {
	if config == nil {
		return defaultTimeout * time.Second
	} else {
		return time.Duration(config.Timeout) * time.Second
	}
}

func getEndpoint(config *ClientConfig) string {
	if config == nil {
		return defaultEndpoint
	} else {
		return config.GetConnectionString()
	}
}

func checkError(err error) error {
	if err != nil {
		return err
	} else {
		return nil
	}
}
