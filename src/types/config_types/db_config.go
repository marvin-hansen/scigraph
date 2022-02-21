// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package config_types

type DBConfig struct {
	Prod             bool
	SvcHost          string
	SvcVersion       string
	DBConfigName     string
	DBLocalUri       string
	DBClusterURI     string
	DBClusterUserENV string
	DBClusterPWENV   string
	DBPort           string
	DBUserLocal      string
	DBPasswordLocal  string
	DBName           string
	DBSSL            string
	DBModel          []interface{}
	DBCompositeTypes []interface{}
}
