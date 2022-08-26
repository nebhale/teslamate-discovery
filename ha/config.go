// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package ha

const (
	DefaultDiscoveryPrefix = "homeassistant"
)

var DefaultConfig = Config{
	DiscoveryPrefix: DefaultDiscoveryPrefix,
}

type Config struct {
	DiscoveryPrefix string `mapstructure:"discovery_prefix"`
}
