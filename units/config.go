// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package units

const (
	DefaultDistance  = Imperial
	DefaultPressure  = Imperial
	DefaultRangeType = Rated
)

var DefaultConfig = Config{
	Distance:  DefaultDistance,
	Pressure:  DefaultPressure,
	RangeType: DefaultRangeType,
}

type Config struct {
	Distance  SystemOfMeasurement `mapstructure:"distance"`
	Pressure  SystemOfMeasurement `mapstructure:"pressure"`
	RangeType RangeType           `mapstructure:"range_type"`
}
