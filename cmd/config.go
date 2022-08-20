// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nebhale/teslamate-discovery/ha"
	"github.com/nebhale/teslamate-discovery/mqtt"
	"github.com/nebhale/teslamate-discovery/tm"
	"github.com/nebhale/teslamate-discovery/units"
)

var DefaultConfig = Config{
	HomeAssistant: ha.DefaultConfig,
	MQTT:          mqtt.DefaultConfig,
	Teslamate:     tm.DefaultConfig,
	Units:         units.DefaultConfig,
}

type Config struct {
	HomeAssistant ha.Config    `mapstructure:"ha"`
	MQTT          mqtt.Config  `mapstructure:"mqtt"`
	Teslamate     tm.Config    `mapstructure:"tm"`
	Units         units.Config `mapstructure:"units"`
}

func UnmarshalConfig(config *Config, v *viper.Viper) CobraEFn {
	return func(cmd *cobra.Command, args []string) error {
		v.SetConfigName("teslamate-discovery")
		v.SetConfigType("yaml")

		dir, err := os.UserConfigDir()
		if err == nil {
			v.AddConfigPath(fmt.Sprintf("%s/teslamate-discovery", dir))
		}
		v.AddConfigPath("$HOME")
		v.AddConfigPath(".")

		if err := v.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return err
			}
		}

		if err := v.Unmarshal(&config); err != nil {
			return err
		}

		if config.MQTT.Username == "" {
			return fmt.Errorf("mqtt username must be specified")
		}

		if config.MQTT.Password == "" {
			return fmt.Errorf("mqtt password must be specified")
		}

		return nil
	}
}
