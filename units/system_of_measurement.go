// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package units

import (
	"fmt"

	"github.com/spf13/cobra"
)

const RoundingValueTemplate = "{{ value | round(1) }}"

type SystemOfMeasurement string

const (
	Imperial SystemOfMeasurement = "imperial"
	Metric   SystemOfMeasurement = "metric"
)

func (s SystemOfMeasurement) DistanceLongUnits() string {
	if s == Metric {
		return "km"
	}

	return "mi"
}

func (s SystemOfMeasurement) DistanceLongValueTemplate() string {
	if s == Metric {
		return RoundingValueTemplate
	}

	return "{{ (value | float(0) / 1.609344) | round(1) }}"
}

func (s SystemOfMeasurement) DistanceShortUnits() string {
	if s == Metric {
		return "m"
	}

	return "ft"
}

func (s SystemOfMeasurement) DistanceShortValueTemplate() string {
	if s == Metric {
		return RoundingValueTemplate
	}

	return "{{ (value | float(0) * 3.280839) | round(1) }}"
}

func (s SystemOfMeasurement) PressureUnits() string {
	if s == Metric {
		return "bar"
	}

	return "psi"
}

func (s SystemOfMeasurement) PressureValueTemplate() string {
	if s == Metric {
		return RoundingValueTemplate
	}

	return "{{ (value | float(0) * 14.503773) | round(1) }}"
}

func (s *SystemOfMeasurement) Set(v string) error {
	switch v {
	case "imperial", "metric":
		*s = SystemOfMeasurement(v)
	default:
		return fmt.Errorf("must be one of imperial, metric")
	}
	return nil
}

func (s SystemOfMeasurement) String() string {
	return string(s)
}

func (s SystemOfMeasurement) SpeedUnits() string {
	if s == Metric {
		return "kph"
	}

	return "mph"
}

func (s SystemOfMeasurement) SpeedValueTemplate() string {
	return s.DistanceLongValueTemplate()
}

func (s SystemOfMeasurement) Type() string {
	return "string"
}

func SystemOfMeasurementCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{string(Imperial), string(Metric)}, cobra.ShellCompDirectiveDefault
}
