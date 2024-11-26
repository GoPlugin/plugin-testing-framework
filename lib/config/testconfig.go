package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"

	"github.com/goplugin/plugin-testing-framework/seth"

	"github.com/goplugin/plugin-testing-framework/lib/logging"
)

func (c *TestConfig) GetLoggingConfig() *LoggingConfig {
	return c.Logging
}

func (c *TestConfig) GetNodeConfig() *NodeConfig {
	return c.NodeConfig
}

func (c TestConfig) GetNetworkConfig() *NetworkConfig {
	return c.Network
}

func (c TestConfig) GetPluginImageConfig() *PluginImageConfig {
	return c.PluginImage
}

func (c TestConfig) GetPrivateEthereumNetworkConfig() *EthereumNetworkConfig {
	return c.PrivateEthereumNetwork
}

func (c TestConfig) GetPyroscopeConfig() *PyroscopeConfig {
	return c.Pyroscope
}

type TestConfig struct {
	PluginImage         *PluginImageConfig  `toml:"PluginImage"`
	PluginUpgradeImage  *PluginImageConfig  `toml:"PluginUpgradeImage"`
	Logging                *LoggingConfig         `toml:"Logging"`
	Network                *NetworkConfig         `toml:"Network"`
	Pyroscope              *PyroscopeConfig       `toml:"Pyroscope"`
	PrivateEthereumNetwork *EthereumNetworkConfig `toml:"PrivateEthereumNetwork"`
	WaspConfig             *WaspAutoBuildConfig   `toml:"WaspAutoBuild"`
	Seth                   *seth.Config           `toml:"Seth"`
	NodeConfig             *NodeConfig            `toml:"NodeConfig"`
}

// Read config values from environment variables
func (c *TestConfig) ReadFromEnvVar() error {
	logger := logging.GetTestLogger(nil)

	testLogCollect := MustReadEnvVar_Boolean(E2E_TEST_LOG_COLLECT_ENV)
	if testLogCollect != nil {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		logger.Debug().Msgf("Using %s env var to override Logging.TestLogCollect", E2E_TEST_LOG_COLLECT_ENV)
		c.Logging.TestLogCollect = testLogCollect
	}

	loggingRunID := MustReadEnvVar_String(E2E_TEST_LOGGING_RUN_ID_ENV)
	if loggingRunID != "" {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		logger.Debug().Msgf("Using %s env var to override Logging.RunID", E2E_TEST_LOGGING_RUN_ID_ENV)
		c.Logging.RunId = &loggingRunID
	}

	logstreamLogTargets := MustReadEnvVar_Strings(E2E_TEST_LOG_STREAM_LOG_TARGETS_ENV, ",")
	if len(logstreamLogTargets) > 0 {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		if c.Logging.LogStream == nil {
			c.Logging.LogStream = &LogStreamConfig{}
		}
		logger.Debug().Msgf("Using %s env var to override Logging.LogStream.LogTargets", E2E_TEST_LOG_STREAM_LOG_TARGETS_ENV)
		c.Logging.LogStream.LogTargets = logstreamLogTargets
	}

	lokiTenantID := MustReadEnvVar_String(E2E_TEST_LOKI_TENANT_ID_ENV)
	if lokiTenantID != "" {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		if c.Logging.Loki == nil {
			c.Logging.Loki = &LokiConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Logging.Loki.TenantId", E2E_TEST_LOKI_TENANT_ID_ENV)
		c.Logging.Loki.TenantId = &lokiTenantID
	}

	lokiEndpoint := MustReadEnvVar_String(E2E_TEST_LOKI_ENDPOINT_ENV)
	if lokiEndpoint != "" {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		if c.Logging.Loki == nil {
			c.Logging.Loki = &LokiConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Logging.Loki.Endpoint", E2E_TEST_LOKI_ENDPOINT_ENV)
		c.Logging.Loki.Endpoint = &lokiEndpoint
	}

	lokiBasicAuth := MustReadEnvVar_String(E2E_TEST_LOKI_BASIC_AUTH_ENV)
	if lokiBasicAuth != "" {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		if c.Logging.Loki == nil {
			c.Logging.Loki = &LokiConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Logging.Loki.BasicAuth", E2E_TEST_LOKI_BASIC_AUTH_ENV)
		c.Logging.Loki.BasicAuth = &lokiBasicAuth
	}

	lokiBearerToken := MustReadEnvVar_String(E2E_TEST_LOKI_BEARER_TOKEN_ENV)
	if lokiBearerToken != "" {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		if c.Logging.Loki == nil {
			c.Logging.Loki = &LokiConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Logging.Loki.BearerToken", E2E_TEST_LOKI_BEARER_TOKEN_ENV)
		c.Logging.Loki.BearerToken = &lokiBearerToken
	}

	grafanaBaseUrl := MustReadEnvVar_String(E2E_TEST_GRAFANA_BASE_URL_ENV)
	if grafanaBaseUrl != "" {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		if c.Logging.Grafana == nil {
			c.Logging.Grafana = &GrafanaConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Logging.Grafana.BaseUrl", E2E_TEST_GRAFANA_BASE_URL_ENV)
		c.Logging.Grafana.BaseUrl = &grafanaBaseUrl
	}

	grafanaDashboardUrl := MustReadEnvVar_String(E2E_TEST_GRAFANA_DASHBOARD_URL_ENV)
	if grafanaDashboardUrl != "" {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		if c.Logging.Grafana == nil {
			c.Logging.Grafana = &GrafanaConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Logging.Grafana.DashboardUrl", E2E_TEST_GRAFANA_DASHBOARD_URL_ENV)
		c.Logging.Grafana.DashboardUrl = &grafanaDashboardUrl
	}

	grafanaBearerToken := MustReadEnvVar_String(E2E_TEST_GRAFANA_BEARER_TOKEN_ENV)
	if grafanaBearerToken != "" {
		if c.Logging == nil {
			c.Logging = &LoggingConfig{}
		}
		if c.Logging.Grafana == nil {
			c.Logging.Grafana = &GrafanaConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Logging.Grafana.BearerToken", E2E_TEST_GRAFANA_BEARER_TOKEN_ENV)
		c.Logging.Grafana.BearerToken = &grafanaBearerToken
	}

	pyroscopeEnabled := MustReadEnvVar_Boolean(E2E_TEST_PYROSCOPE_ENABLED_ENV)
	if pyroscopeEnabled != nil {
		if c.Pyroscope == nil {
			c.Pyroscope = &PyroscopeConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Pyroscope.Enabled", E2E_TEST_PYROSCOPE_ENABLED_ENV)
		c.Pyroscope.Enabled = pyroscopeEnabled
	}

	pyroscopeServerUrl := MustReadEnvVar_String(E2E_TEST_PYROSCOPE_SERVER_URL_ENV)
	if pyroscopeServerUrl != "" {
		if c.Pyroscope == nil {
			c.Pyroscope = &PyroscopeConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Pyroscope.ServerUrl", E2E_TEST_PYROSCOPE_SERVER_URL_ENV)
		c.Pyroscope.ServerUrl = &pyroscopeServerUrl
	}

	pyroscopeKey := MustReadEnvVar_String(E2E_TEST_PYROSCOPE_KEY_ENV)
	if pyroscopeKey != "" {
		if c.Pyroscope == nil {
			c.Pyroscope = &PyroscopeConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Pyroscope.Key", E2E_TEST_PYROSCOPE_KEY_ENV)
		c.Pyroscope.Key = &pyroscopeKey
	}

	pyroscopeEnvironment := MustReadEnvVar_String(E2E_TEST_PYROSCOPE_ENVIRONMENT_ENV)
	if pyroscopeEnvironment != "" {
		if c.Pyroscope == nil {
			c.Pyroscope = &PyroscopeConfig{}
		}
		logger.Info().Msgf("Using %s env var to override Pyroscope.Environment", E2E_TEST_PYROSCOPE_ENVIRONMENT_ENV)
		c.Pyroscope.Environment = &pyroscopeEnvironment
	}

	selectedNetworks := MustReadEnvVar_Strings(E2E_TEST_SELECTED_NETWORK_ENV, ",")
	if len(selectedNetworks) > 0 {
		if c.Network == nil {
			c.Network = &NetworkConfig{}
		}
		logger.Debug().Msgf("Using %s env var to override Network.SelectedNetworks", E2E_TEST_SELECTED_NETWORK_ENV)
		c.Network.SelectedNetworks = selectedNetworks
	}

	walletKeys := ReadEnvVarGroupedMap(E2E_TEST_WALLET_KEY_ENV, E2E_TEST_WALLET_KEYS_ENV)
	if len(walletKeys) > 0 {
		if c.Network == nil {
			c.Network = &NetworkConfig{}
		}
		logger.Info().Msgf("Using %s and/or %s env vars to override Network.WalletKeys", E2E_TEST_WALLET_KEY_ENV, E2E_TEST_WALLET_KEYS_ENV)
		c.Network.WalletKeys = walletKeys
	}

	rpcHttpUrls := ReadEnvVarGroupedMap(E2E_TEST_RPC_HTTP_URL_ENV, E2E_TEST_RPC_HTTP_URLS_ENV)
	if len(rpcHttpUrls) > 0 {
		if c.Network == nil {
			c.Network = &NetworkConfig{}
		}
		logger.Info().Msgf("Using %s and/or %s env vars to override Network.RpcHttpUrls", E2E_TEST_RPC_HTTP_URL_ENV, E2E_TEST_RPC_HTTP_URLS_ENV)
		c.Network.RpcHttpUrls = rpcHttpUrls
	}

	rpcWsUrls := ReadEnvVarGroupedMap(E2E_TEST_RPC_WS_URL_ENV, E2E_TEST_RPC_WS_URLS_ENV)
	if len(rpcWsUrls) > 0 {
		if c.Network == nil {
			c.Network = &NetworkConfig{}
		}
		logger.Info().Msgf("Using %s and/or %s env vars to override Network.RpcWsUrls", E2E_TEST_RPC_WS_URL_ENV, E2E_TEST_RPC_WS_URLS_ENV)
		c.Network.RpcWsUrls = rpcWsUrls
	}

	pluginImage := MustReadEnvVar_String(E2E_TEST_PLUGIN_IMAGE_ENV)
	if pluginImage != "" {
		if c.PluginImage == nil {
			c.PluginImage = &PluginImageConfig{}
		}
		logger.Info().Msgf("Using %s env var to override PluginImage.Image", E2E_TEST_PLUGIN_IMAGE_ENV)
		c.PluginImage.Image = &pluginImage
	}

	pluginVersion := MustReadEnvVar_String(E2E_TEST_PLUGIN_VERSION_ENV)
	if pluginVersion != "" {
		if c.PluginImage == nil {
			c.PluginImage = &PluginImageConfig{}
		}
		logger.Debug().Msgf("Using %s env var to override PluginImage.Version", E2E_TEST_PLUGIN_VERSION_ENV)
		c.PluginImage.Version = &pluginVersion
	}

	pluginPostgresVersion := MustReadEnvVar_String(E2E_TEST_PLUGIN_POSTGRES_VERSION_ENV)
	if pluginPostgresVersion != "" {
		if c.PluginImage == nil {
			c.PluginImage = &PluginImageConfig{}
		}
		logger.Debug().Msgf("Using %s env var to override PluginImage.PostgresVersion", E2E_TEST_PLUGIN_POSTGRES_VERSION_ENV)
		c.PluginImage.PostgresVersion = &pluginPostgresVersion
	}

	pluginUpgradeImage := MustReadEnvVar_String(E2E_TEST_PLUGIN_UPGRADE_IMAGE_ENV)
	if pluginUpgradeImage != "" {
		if c.PluginUpgradeImage == nil {
			c.PluginUpgradeImage = &PluginImageConfig{}
		}
		logger.Info().Msgf("Using %s env var to override PluginUpgradeImage.Image", E2E_TEST_PLUGIN_UPGRADE_IMAGE_ENV)
		c.PluginUpgradeImage.Image = &pluginUpgradeImage
	}

	pluginUpgradeVersion := MustReadEnvVar_String(E2E_TEST_PLUGIN_UPGRADE_VERSION_ENV)
	if pluginUpgradeVersion != "" {
		if c.PluginUpgradeImage == nil {
			c.PluginUpgradeImage = &PluginImageConfig{}
		}
		logger.Debug().Msgf("Using %s env var to override PluginUpgradeImage.Version", E2E_TEST_PLUGIN_UPGRADE_VERSION_ENV)
		c.PluginUpgradeImage.Version = &pluginUpgradeVersion
	}

	return nil
}

func LoadSecretEnvsFromFiles() error {
	logger := logging.GetTestLogger(nil)

	// Load existing environment variables into a map
	existingEnv := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		existingEnv[pair[0]] = pair[1]
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrapf(err, "error getting user home directory")
	}
	homePath := fmt.Sprintf("%s/.testsecrets", homeDir)
	etcPath := "/etc/e2etests/.testsecrets"
	testsecretsPath := []string{etcPath, homePath}

	for _, path := range testsecretsPath {
		logger.Debug().Msgf("Checking for test secrets file at %s", path)

		// Load variables from the env file
		envMap, err := godotenv.Read(path)
		if err != nil {
			if os.IsNotExist(err) {
				logger.Debug().Msgf("No test secrets file found at %s", path)
				continue
			}
			return fmt.Errorf("error reading test secrets file at %s", path)
		}

		// Set env vars from file only if they are not already set
		for key, value := range envMap {
			if _, exists := existingEnv[key]; !exists {
				logger.Debug().Msgf("Setting env var %s from %s file", key, path)
				os.Setenv(key, value)
			} else {
				logger.Debug().Msgf("Env var %s already set, not overriding it from %s file", key, path)
			}
		}
	}

	return nil
}
