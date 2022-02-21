// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package config_types

import "fmt"

type ImxConfig struct {
	IbAccount string
	ClientID  string
}

func (ibc ImxConfig) String() string {
	return fmt.Sprintf("ImxConfig<IbAccount: %v, ClientID: %v>",
		ibc.IbAccount,
		ibc.ClientID,
	)
}
