package config

import "github.com/goplugin/plugin-testing-framework/seth"

type SethConfig interface {
	GetSethConfig() *seth.Config
}

type NamedConfigurations interface {
	GetConfigurationNames() []string
}

type GlobalTestConfig interface {
	GetPluginImageConfig() *PluginImageConfig
	GetLoggingConfig() *LoggingConfig
	GetNetworkConfig() *NetworkConfig
	GetPrivateEthereumNetworkConfig() *EthereumNetworkConfig
	GetPyroscopeConfig() *PyroscopeConfig
	GetNodeConfig() *NodeConfig
	SethConfig
}
