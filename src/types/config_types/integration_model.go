package config_types

import (
	"fmt"
)

func NewExchangeIntegration(exchangeName ExchangeName, integrationName string,
	integrationVersion string, integrationChannels ChannelConfig,
) *ExchangeIntegration {
	return &ExchangeIntegration{
		ExchangeName:        exchangeName,
		IntegrationName:     integrationName,
		IntegrationVersion:  integrationVersion,
		IntegrationHandle:   integrationName + integrationVersion,
		IntegrationChannels: integrationChannels,
	}
}

type ExchangeIntegration struct {
	// tableName is an optional field that specifies custom table name and alias.
	// By default, go-pg generates table name and alias from struct name.
	tableName           struct{} `pg:"imdb.exchange_integration_config,alias:integration_config"` // define DB schema in dot notation db_schema.table_name
	id                  int64    `pg:",pk,unique"`                                                // PK for internal DB use
	ExchangeName        ExchangeName
	IntegrationName     string
	IntegrationVersion  string
	IntegrationHandle   string // name+version
	IntegrationChannels ChannelConfig
}

func (c ExchangeIntegration) String() string {
	return fmt.Sprintf("[ExchangeIntegration]:  ExchangeName: %v, IntegrationName: %v, IntegrationVersion: %v, IntegrationHandle: %v, IntegrationChannels: %v",
		c.ExchangeName,
		c.IntegrationName,
		c.IntegrationVersion,
		c.IntegrationHandle,
		c.IntegrationChannels,
	)
}

func (c ExchangeIntegration) GetIntegrationHandle() string {
	return c.IntegrationName + c.IntegrationVersion
}
