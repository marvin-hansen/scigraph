// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package config_types

type MainConfig struct {
	//DBConf      *pgdb_client.DBConfig
	ServiceID   string
	ExchangeID  string // for IMX services that integrate an exchange
	ServiceName string
	Port        string
	Network     string
}
