package config

import (
	"errors"
)

type PluginImageConfig struct {
	Image           *string `toml:"-"`
	Version         *string `toml:"version"`
	PostgresVersion *string `toml:"postgres_version,omitempty"`
}

// Validate checks that the plugin image config is valid, which means that
// both image and version are set and non-empty
func (c *PluginImageConfig) Validate() error {
	if c.Image == nil || *c.Image == "" {
		return errors.New("plugin image name must be set")
	}

	if c.Version == nil || *c.Version == "" {
		return errors.New("plugin version must be set")
	}

	return nil
}
