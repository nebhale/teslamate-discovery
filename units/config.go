// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package units

const (
	DefaultDistance = Imperial
	DefaultPressure = Imperial
)

var DefaultConfig = Config{
	Distance: DefaultDistance,
	Pressure: DefaultPressure,
}

type Config struct {
	Distance SystemOfMeasurement `mapstructure:"distance"`
	Pressure SystemOfMeasurement `mapstructure:"pressure"`
}
