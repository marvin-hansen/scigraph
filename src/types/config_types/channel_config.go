// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package config_types

import (
	"fmt"
	"strings"
)

type ChannelConfigType string

func GetImxChannelConfig(exchangeID ExchangeName) ChannelConfig {
	clientID := strings.ToLower(exchangeID.String())
	channelArray := GetChannelArray()
	return ChannelConfig{
		ClientID:       clientID,
		AccountChannel: channelArray[0],
		ApiChannel:     channelArray[1],
		ControlChannel: channelArray[2],
		DataChannel:    channelArray[3],
		OrderChannel:   channelArray[4],
	}
}

func GetClientChannelConfig(clientID string) ChannelConfig {
	channelArray := GetChannelArray()
	return ChannelConfig{
		ClientID:       clientID,
		AccountChannel: channelArray[0],
		ApiChannel:     channelArray[1],
		ControlChannel: channelArray[2],
		DataChannel:    channelArray[3],
		OrderChannel:   channelArray[4],
	}
}

// GetChannelArray returns all exchange integration channels
// Size is 15 because of the max size of 15 channels allows in model.Service config
func GetChannelArray() [15]string {
	var chArray [15]string
	chArray[0] = "account"
	chArray[1] = "api"
	chArray[2] = "control"
	chArray[3] = "data"
	chArray[4] = "order"
	return chArray
}

type ChannelConfig struct {
	Id             int64 `pg:",pk,unique"` // PK for internal DB use
	ClientID       string
	AccountChannel string
	ApiChannel     string
	ControlChannel string
	DataChannel    string
	OrderChannel   string
}

func (c ChannelConfig) String() string {
	return fmt.Sprintf("[ChannelConfig]: ClientID: %v, AccountChannel: %v, ApiChannel: %v, ControlChannel: %v, DataChannel: %v, OrderChannel: %v",
		c.ClientID,
		c.AccountChannel,
		c.ApiChannel,
		c.ControlChannel,
		c.DataChannel,
		c.OrderChannel,
	)
}

func (c ChannelConfig) GetClientID() string {
	return c.ClientID
}

func (c ChannelConfig) GetClientAmxAccountChannel() string {
	return c.ClientID + "-" + c.AccountChannel
}

func (c ChannelConfig) GetClientApxApiChannel() string {
	return c.ClientID + "-" + c.ApiChannel
}

func (c ChannelConfig) GetClientCmxControlChannel() string {
	return c.ClientID + "-" + c.ControlChannel
}

func (c ChannelConfig) GetClientDmxDataChannel() string {
	return c.ClientID + "-" + c.DataChannel
}

func (c ChannelConfig) GetClientEmxOrderExecutionChannel() string {
	return c.ClientID + "-" + c.OrderChannel
}
