package main

import (
	ctf_config "github.com/goplugin/plugin-testing-framework/lib/config"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/environment"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/plugin"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/ethereum"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/mockserver"
	mockservercfg "github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/mockserver-cfg"
	"github.com/goplugin/plugin-testing-framework/lib/utils/ptr"
)

func main() {
	// in actual implementation here you should read the config from TOML file instead of creating structs manually
	pluginConfig := ctf_config.PluginImageConfig{
		Image:           ptr.Ptr("public.ecr.aws/plugin/plugin"),
		Version:         ptr.Ptr("v2.8.0"),
		PostgresVersion: ptr.Ptr("15.6"),
	}

	pyroscope := ctf_config.PyroscopeConfig{
		Enabled: ptr.Ptr(false),
	}

	config := struct {
		Plugin ctf_config.PluginImageConfig
		Pyroscope ctf_config.PyroscopeConfig
	}{
		Plugin: pluginConfig,
		Pyroscope: pyroscope,
	}

	var overrideFn = func(_ interface{}, target interface{}) {
		ctf_config.MustConfigOverridePluginVersion(&pluginConfig, target)
		ctf_config.MightConfigOverridePyroscopeKey(&pyroscope, target)
	}

	err := environment.New(&environment.Config{
		NamespacePrefix:   "ztest",
		KeepConnection:    true,
		RemoveOnInterrupt: true,
	}).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.NewWithOverride(0, map[string]interface{}{
			"replicas": 1,
		}, &config, overrideFn)).
		Run()
	if err != nil {
		panic(err)
	}
}
