// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package config_types

import "scigraph/src/clients/pgdb_client"

type ServiceConfig struct {
	AMDBHost   string
	CMDBHost   string
	DBHost     string
	DBConfig   pgdb_client.DBConfig
	ImdbHost   string
	NatsHost   string
	NatsConfig ChannelConfig
	ImxConfig  ImxConfig
}
