// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package config_types

import "fmt"

type Service struct {
	// tableName is an optional field that specifies custom table name and alias.
	// By default go-pg generates table name and alias from struct name.
	tableName                   struct{} `pg:"smdb.service,alias:svc"` // default values are the same
	Id                          int64    `pg:",pk,unique"`             // PK for internal DB use
	ServiceHandle               string   // Convenience accessor: name + version
	ServiceName                 string
	ServiceVersion              string
	ServiceDescription          string
	ServiceRegistered           string
	ServiceHealthUri            string
	ServiceBaseUri              string
	ServiceExposure             ServiceExposureType   // CmxControlChannel flow flag for auto-config.
	ServiceEndpoint             Endpoint              // Only externally developed services (i.e. DB) expose usually one endpoint.
	ServiceChannels             [15]string            `pg:",array"` // PG array stuffs everything into one single fat service table. No relations.
	ServiceDependencies         []Dependency          `pg:",array"`
	ServiceExchangeIntegrations []ExchangeIntegration `pg:",array"`
}

func (s Service) String() string {
	return fmt.Sprintf("[Service]: Id: %v, ServiceHandle: %v, ServiceName: %v, ServiceVersion: %v, ServiceDescription: %v, ServiceRegistered: %v, ServiceHealthUri: %v, ServiceBaseUri: %v, ServiceDependencies: %v, ServiceEndpoints: %v",
		s.Id,
		s.ServiceHandle,
		s.ServiceName,
		s.ServiceVersion,
		s.ServiceDescription,
		s.ServiceRegistered,
		s.ServiceHealthUri,
		s.ServiceBaseUri,
		s.ServiceDependencies,
		s.ServiceEndpoint)
}

type Dependency struct {
	Id                    int64 `pg:",pk,unique"` // PK for internal DB use
	DependencyName        string
	DependencyVersion     string
	DependencyDescription string
}

func (s Dependency) String() string {
	return fmt.Sprintf("[Dependency]: Id: %v, DependencyName: %v, DependencyVersion: %v, DependencyDescription: %v",
		s.Id,
		s.DependencyName,
		s.DependencyVersion,
		s.DependencyDescription)
}

type Endpoint struct {
	Id                  int64 `pg:",pk,unique"` // PK for internal DB use
	EndpointName        string
	EndpointVersion     string
	EndpointPort        string
	EndpointUri         string
	EndpointProtocol    string
	EndpointEncoding    string
	EndpointDescription string
}

func (s Endpoint) String() string {
	return fmt.Sprintf("[Endpoint]: Id: %v, EndpointName: %v, EndpointVersion: %v, EndpointPort: %v, EndpointUri: %v, EndpointProtocol: %v, EndpointEncoding: %v, EndpointDescription: %v",
		s.Id,
		s.EndpointName,
		s.EndpointVersion,
		s.EndpointPort,
		s.EndpointUri,
		s.EndpointProtocol,
		s.EndpointEncoding,
		s.EndpointDescription)
}

type HostEndpoint struct {
	HostUri      string
	EndpointPort string
}

func (s HostEndpoint) String() string {
	return fmt.Sprintf("[HostEndpoint]: HostUri: %v, EndpointPort: %v",
		s.HostUri,
		s.HostUri)
}
