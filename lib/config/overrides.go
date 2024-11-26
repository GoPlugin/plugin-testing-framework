package config

import (
	"dario.cat/mergo"
)

// MustConfigOverridePluginVersion will override the plugin image and version
// in property maps from the passed config. It will panic if the config is nil.
func MustConfigOverridePluginVersion(config *PluginImageConfig, target interface{}) {
	if config == nil {
		panic("[PluginImageConfig] must be present")
	}
	if config.Image != nil && *config.Image != "" && config.Version != nil && *config.Version != "" {
		if err := mergo.Merge(target, map[string]interface{}{
			"plugin": map[string]interface{}{
				"image": map[string]interface{}{
					"image":   *config.Image,
					"version": *config.Version,
				},
			},
		}, mergo.WithOverride); err != nil {
			panic(err)
		}
	}

	if config.PostgresVersion != nil && *config.PostgresVersion != "" {
		if err := mergo.Merge(target, map[string]interface{}{
			"db": map[string]interface{}{
				"image": map[string]interface{}{
					"version": *config.PostgresVersion,
				},
			},
		}, mergo.WithOverride); err != nil {
			panic(err)
		}
	}
}

// MightConfigOverridePyroscope will override the pyroscope config in property maps
// from the passed config. If the config is nil, or the enabled flag is not set, or
// the key is not set, then this function will do nothing.
func MightConfigOverridePyroscopeKey(config *PyroscopeConfig, target interface{}) {
	if config == nil || (config.Enabled == nil || !*config.Enabled) || (config.Key == nil || *config.Key == "") {
		return
	}

	env := make(map[string]string)
	env["CL_PYROSCOPE_AUTH_TOKEN"] = *config.Key

	if err := mergo.Merge(target, map[string]interface{}{
		"env": env,
	}, mergo.WithOverride); err != nil {
		panic(err)
	}
}
