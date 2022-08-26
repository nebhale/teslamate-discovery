// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package units_test

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"

	. "github.com/nebhale/teslamate-discovery/units"
)

func TestSystemOfMeasurement_DistanceLongUnits(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "mi",
		},
		{
			name: "metric",
			s:    Metric,
			want: "km",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DistanceLongUnits(); got != tt.want {
				t.Errorf("SystemOfMeasurement.DistanceLongUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_DistanceLongValueTemplate(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "{{ (value | float(0) / 1.609344) | round(1) }}",
		},
		{
			name: "metric",
			s:    Metric,
			want: "{{ value | round(1) }}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DistanceLongValueTemplate(); got != tt.want {
				t.Errorf("SystemOfMeasurement.DistanceLongValueTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_DistanceShortUnits(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "ft",
		},
		{
			name: "metric",
			s:    Metric,
			want: "m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DistanceShortUnits(); got != tt.want {
				t.Errorf("SystemOfMeasurement.DistanceShortUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_DistanceShortValueTemplate(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "{{ (value | float(0) * 3.280839) | round(1) }}",
		},
		{
			name: "metric",
			s:    Metric,
			want: "{{ value | round(1) }}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DistanceShortValueTemplate(); got != tt.want {
				t.Errorf("SystemOfMeasurement.DistanceShortValueTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_PressureUnits(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "psi",
		},
		{
			name: "metric",
			s:    Metric,
			want: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.PressureUnits(); got != tt.want {
				t.Errorf("SystemOfMeasurement.PressureUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_PressureValueTemplate(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "{{ (value | float(0) * 14.503773) | round(1) }}",
		},
		{
			name: "metric",
			s:    Metric,
			want: "{{ value | round(1) }}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.PressureValueTemplate(); got != tt.want {
				t.Errorf("SystemOfMeasurement.PressureValueTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_Set(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    SystemOfMeasurement
		wantErr bool
	}{
		{
			name: "imperial",
			args: args{v: "imperial"},
			want: Imperial,
		},
		{
			name: "metric",
			args: args{v: "metric"},
			want: Metric,
		},
		{
			name:    "unknown",
			args:    args{v: "unknown"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SystemOfMeasurement
			err := s.Set(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("SystemOfMeasurement.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if s != tt.want {
				t.Errorf("SystemOfMeasurement.Set() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_String(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "imperial",
		},
		{
			name: "metric",
			s:    Metric,
			want: "metric",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("SystemOfMeasurement.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_SpeedUnits(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "mph",
		},
		{
			name: "metric",
			s:    Metric,
			want: "kph",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SpeedUnits(); got != tt.want {
				t.Errorf("SystemOfMeasurement.SpeedUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_SpeedValueTemplate(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "{{ (value | float(0) / 1.609344) | round(1) }}",
		},
		{
			name: "metric",
			s:    Metric,
			want: "{{ value | round(1) }}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SpeedValueTemplate(); got != tt.want {
				t.Errorf("SystemOfMeasurement.SpeedValueTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurement_Type(t *testing.T) {
	tests := []struct {
		name string
		s    SystemOfMeasurement
		want string
	}{
		{
			name: "imperial",
			s:    Imperial,
			want: "string",
		},
		{
			name: "metric",
			s:    Metric,
			want: "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); got != tt.want {
				t.Errorf("SystemOfMeasurement.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemOfMeasurementCompletion(t *testing.T) {
	type args struct {
		cmd        *cobra.Command
		args       []string
		toComplete string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 cobra.ShellCompDirective
	}{
		{
			name:  "always",
			args:  args{},
			want:  []string{"imperial", "metric"},
			want1: cobra.ShellCompDirectiveDefault,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SystemOfMeasurementCompletion(tt.args.cmd, tt.args.args, tt.args.toComplete)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemOfMeasurementCompletion() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SystemOfMeasurementCompletion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
