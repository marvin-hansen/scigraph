// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package config_types

type EnvironmentType uint8

const (
	UNKNOWN_ENV EnvironmentType = iota // 0 :
	LOCAL
	CLUSTER
)

func (e EnvironmentType) String() string {
	return [...]string{"UNKNOWN_ENV", "LOCAL", "CLUSTER"}[e]
}
