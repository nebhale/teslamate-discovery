// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package mqtt

const (
	DefaultScheme = "ssl"
	DefaultHost   = "127.0.0.1"
	DefaultPort   = 8883
)

var DefaultConfig = Config{
	Scheme: DefaultScheme,
	Host:   DefaultHost,
	Port:   DefaultPort,
}

type Config struct {
	Scheme   string `mapstructure:"scheme"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
