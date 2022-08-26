// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nebhale/teslamate-discovery/ha"
	"github.com/nebhale/teslamate-discovery/mqtt"
	"github.com/nebhale/teslamate-discovery/tm"
	"github.com/nebhale/teslamate-discovery/units"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	config := &DefaultConfig

	cmd, viper := CreateCommand()
	cmd.PreRunE = UnmarshalConfig(config, viper)
	cmd.RunE = Run(config)

	if err := cmd.ExecuteContext(ctx); err != nil {
		stop()
		os.Exit(1)
	}

	stop()
}

type CobraEFn func(cmd *cobra.Command, args []string) error

func CreateCommand() (*cobra.Command, *viper.Viper) {
	viper := viper.New()

	cmd := &cobra.Command{
		Use:     "teslamate-discovery",
		Short:   "Configure Home Assistant MQTT Discovery for TeslaMate instances",
		Version: Version,
	}

	flags := cmd.Flags()

	_ = flags.String("ha-discovery-prefix", ha.DefaultDiscoveryPrefix, "home assistant discovery message prefix")
	_ = viper.BindPFlag("ha.discovery_prefix", flags.Lookup("ha-discovery-prefix"))
	_ = viper.BindEnv("ha.discovery_prefix", "HA_DISCOVERY_PREFIX")
	viper.SetDefault("ha.discovery_prefix", ha.DefaultDiscoveryPrefix)

	_ = flags.StringP("mqtt-scheme", "s", mqtt.DefaultScheme, "mqtt broker scheme")
	_ = viper.BindPFlag("mqtt.scheme", flags.Lookup("mqtt-scheme"))
	_ = viper.BindEnv("mqtt.scheme", "MQTT_SCHEME")
	viper.SetDefault("mqtt.scheme", mqtt.DefaultScheme)

	_ = flags.StringP("mqtt-host", "h", mqtt.DefaultHost, "mqtt broker host")
	_ = viper.BindPFlag("mqtt.host", flags.Lookup("mqtt-host"))
	_ = viper.BindEnv("mqtt.host", "MQTT_HOST")
	viper.SetDefault("mqtt.host", mqtt.DefaultHost)

	_ = flags.IntP("mqtt-port", "p", mqtt.DefaultPort, "mqtt broker port")
	_ = viper.BindPFlag("mqtt.port", flags.Lookup("mqtt-port"))
	_ = viper.BindEnv("mqtt.port", "MQTT_PORT")
	viper.SetDefault("mqtt.port", mqtt.DefaultPort)

	_ = flags.StringP("mqtt-username", "u", "", "mqtt broker username")
	_ = viper.BindPFlag("mqtt.username", flags.Lookup("mqtt-username"))
	_ = viper.BindEnv("mqtt.username", "MQTT_USERNAME")

	_ = flags.StringP("mqtt-password", "P", "", "mqtt broker password")
	_ = viper.BindPFlag("mqtt.password", flags.Lookup("mqtt-password"))
	_ = viper.BindEnv("mqtt.password", "MQTT_PASSWORD")

	_ = flags.String("tm-prefix", tm.DefaultPrefix, "teslamate message prefix")
	_ = viper.BindPFlag("tm.prefix", flags.Lookup("tm-prefix"))
	_ = viper.BindEnv("tm.prefix", "TM_PREFIX")
	viper.SetDefault("tm.prefix", tm.DefaultPrefix)

	d := units.DefaultDistance
	flags.Var(&d, "units-distance", "distance units [\"imperial\", \"metric\"]")
	_ = cmd.RegisterFlagCompletionFunc("units-distance", units.SystemOfMeasurementCompletion)
	_ = viper.BindPFlag("units.distance", flags.Lookup("units-distance"))
	_ = viper.BindEnv("units.distance", "UNITS_DISTANCE")
	viper.SetDefault("units.distance", units.DefaultDistance)

	p := units.DefaultPressure
	flags.Var(&p, "units-pressure", "pressure units [\"imperial\", \"metric\"]")
	_ = cmd.RegisterFlagCompletionFunc("units-pressure", units.SystemOfMeasurementCompletion)
	_ = viper.BindPFlag("units.pressure", flags.Lookup("units-pressure"))
	_ = viper.BindEnv("units.pressure", "UNITS_PRESSURE")
	viper.SetDefault("units.pressure", units.DefaultPressure)

	_ = flags.Bool("help", false, fmt.Sprintf("help for %s", cmd.Name()))

	return cmd, viper
}

func Run(config *Config) CobraEFn {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		mqtt, err := mqtt.NewMQTT(ctx, config.MQTT)
		if err != nil {
			return err
		}

		vehicles, err := mqtt.ListVehicles(ctx, config.Teslamate)
		if err != nil {
			return err
		}

		for id, dev := range vehicles {
			if err := mqtt.PublishDiscovery(ctx, id, dev, config.HomeAssistant, config.Units); err != nil {
				return err
			}
		}

		return nil
	}
}
