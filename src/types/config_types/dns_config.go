package config_types

type DNSConfig struct {
	DnsHostExternal string
	DnsPortExternal string
	DnsHostInternal string
	DnsPortInternal string
}

func GetDefaultDNSConfig() DNSConfig {
	return DNSConfig{
		DnsHostExternal: "8.8.8.8",
		DnsPortExternal: ":53",
		DnsHostInternal: "", // determined dynamically based on context
		DnsPortInternal: ":53",
	}
}
