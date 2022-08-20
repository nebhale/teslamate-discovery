// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package tm

const (
	DefaultPrefix = "teslamate"
)

var DefaultConfig = Config{
	Prefix: DefaultPrefix,
}

type Config struct {
	Prefix string `mapstructure:"prefix"`
}
