// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package mqtt_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	paho "github.com/eclipse/paho.mqtt.golang"

	"github.com/nebhale/teslamate-discovery/ha"
	. "github.com/nebhale/teslamate-discovery/mqtt"
	"github.com/nebhale/teslamate-discovery/units"
)

func TestMQTT_PublishDiscovery(t *testing.T) {
	d := ha.Device{
		Identifiers: []string{"test-id"},
	}

	type fields struct {
		Client stubPubSub
	}
	type args struct {
		ctx      context.Context
		id       string
		device   ha.Device
		haCfg    ha.Config
		unitsCfg units.Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "publish",
			fields: fields{
				Client: stubPubSub{
					publishTokens: []paho.Token{&stubToken{}},
				},
			},
			args: args{
				ctx:      context.Background(),
				id:       "test-id",
				device:   d,
				haCfg:    ha.Config{DiscoveryPrefix: "test-discovery-prefix"},
				unitsCfg: units.Config{},
			},
			want: []string{
				"test-discovery-prefix/sensor/test-id/charge_energy_added/config",
				"test-discovery-prefix/sensor/test-id/limit/config",
				"test-discovery-prefix/sensor/test-id/charger_current/config",
				"test-discovery-prefix/binary_sensor/test-id/charging/config",
				"test-discovery-prefix/binary_sensor/test-id/plug/config",
				"test-discovery-prefix/sensor/test-id/charger_power/config",
				"test-discovery-prefix/sensor/test-id/charger_voltage/config",
				"test-discovery-prefix/sensor/test-id/start_time/config",
				"test-discovery-prefix/sensor/test-id/time_to_charged/config",
				"test-discovery-prefix/sensor/test-id/inside_temp/config",
				"test-discovery-prefix/binary_sensor/test-id/climate/config",
				"test-discovery-prefix/binary_sensor/test-id/preconditioning/config",
				"test-discovery-prefix/sensor/test-id/outside_temp/config",
				"test-discovery-prefix/sensor/test-id/elevation/config",
				"test-discovery-prefix/sensor/test-id/geofence/config",
				"test-discovery-prefix/sensor/test-id/heading/config",
				"test-discovery-prefix/sensor/test-id/latitude/config",
				"test-discovery-prefix/device_tracker/test-id/location/config",
				"test-discovery-prefix/sensor/test-id/longitude/config",
				"test-discovery-prefix/sensor/test-id/power/config",
				"test-discovery-prefix/sensor/test-id/speed/config",
				"test-discovery-prefix/sensor/test-id/battery/config",
				"test-discovery-prefix/binary_sensor/test-id/charge_port/config",
				"test-discovery-prefix/binary_sensor/test-id/doors/config",
				"test-discovery-prefix/binary_sensor/test-id/frunk/config",
				"test-discovery-prefix/binary_sensor/test-id/health/config",
				"test-discovery-prefix/binary_sensor/test-id/locked/config",
				"test-discovery-prefix/binary_sensor/test-id/occupied/config",
				"test-discovery-prefix/sensor/test-id/odometer/config",
				"test-discovery-prefix/sensor/test-id/range/config",
				"test-discovery-prefix/binary_sensor/test-id/sentry_mode/config",
				"test-discovery-prefix/sensor/test-id/state/config",
				"test-discovery-prefix/sensor/test-id/tire_pressure_front_left/config",
				"test-discovery-prefix/sensor/test-id/tire_pressure_front_right/config",
				"test-discovery-prefix/sensor/test-id/tire_pressure_rear_left/config",
				"test-discovery-prefix/sensor/test-id/tire_pressure_rear_right/config",
				"test-discovery-prefix/binary_sensor/test-id/trunk/config",
				"test-discovery-prefix/binary_sensor/test-id/update/config",
				"test-discovery-prefix/binary_sensor/test-id/windows/config",
				"test-discovery-prefix/sensor/test-id/version/config",
			},
		},
		{
			name: "publish error",
			fields: fields{
				Client: stubPubSub{
					publishTokens: []paho.Token{
						&stubToken{err: fmt.Errorf("publish error")},
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				id:       "test-id",
				device:   d,
				haCfg:    ha.Config{DiscoveryPrefix: "test-discovery-prefix"},
				unitsCfg: units.Config{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MQTT{
				Client: &tt.fields.Client,
			}
			if err := m.PublishDiscovery(tt.args.ctx, tt.args.id, tt.args.device, tt.args.haCfg, tt.args.unitsCfg); (err != nil) != tt.wantErr {
				t.Errorf("MQTT.PublishDiscovery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			var topics []string
			for _, v := range tt.fields.Client.publishArgs {
				topics = append(topics, v.topic)
			}

			if !reflect.DeepEqual(topics, tt.want) {
				t.Errorf("MQTT.PublishDiscovery() topics = %v, want %v", topics, tt.want)
				return
			}
		})
	}
}

func TestName(t *testing.T) {
	type args struct {
		device ha.Device
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				device: ha.Device{Name: "test"},
				suffix: "name",
			},
			want: "test name",
		},
		{
			name: "empty",
			args: args{
				device: ha.Device{Name: "test"},
				suffix: "",
			},
			want: "test ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Name(tt.args.device, tt.args.suffix); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStateTopic(t *testing.T) {
	type args struct {
		device ha.Device
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				device: ha.Device{Identifiers: []string{"test"}},
				suffix: "/topic",
			},
			want: "test/topic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StateTopic(tt.args.device, tt.args.suffix); got != tt.want {
				t.Errorf("StateTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueId(t *testing.T) {
	type args struct {
		device ha.Device
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				device: ha.Device{Identifiers: []string{"test"}},
				suffix: "/uniqueId",
			},
			want: "test/uniqueId",
		},
		{
			name: "default",
			args: args{
				device: ha.Device{Identifiers: []string{"test/test"}},
				suffix: "/uniqueId",
			},
			want: "test_test/uniqueId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueId(tt.args.device, tt.args.suffix); got != tt.want {
				t.Errorf("UniqueId() = %v, want %v", got, tt.want)
			}
		})
	}
}
