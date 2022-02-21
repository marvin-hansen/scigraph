// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package config_types

type NatsConfig struct {
	NatsProtocol        string
	NatsHost            string
	NatsPort            string
	NatsDefaultURL      string
	NatsServerUri       string
	MaxReconnects       int
	MaxReconnectWaitSec float64
}

func GetNatsDefaultConfig() NatsConfig {
	return NatsConfig{
		NatsProtocol:        "nats://",
		NatsHost:            "0.0.0.0",
		NatsPort:            ":4222",
		NatsDefaultURL:      "nats://127.0.0.1:4222",
		NatsServerUri:       "", // Set dynamically according to context.
		MaxReconnects:       60,
		MaxReconnectWaitSec: 1,
	}
}
