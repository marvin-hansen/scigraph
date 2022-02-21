// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package config_types

type SMXConfig struct {
	SMXClusterHost string
	SMXClusterPort string
	SMXLocalHost   string
	SMXLocalPort   string
}

func GetSmdbDefaultConfig() SMXConfig {
	return SMXConfig{
		SMXClusterHost: "smdb-service.default.svc.cluster.local",
		SMXClusterPort: ":5050",
		SMXLocalHost:   "localhost",
		SMXLocalPort:   ":5050",
	}
}
