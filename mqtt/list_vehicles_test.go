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
	"github.com/nebhale/teslamate-discovery/tm"
)

func TestMQTT_ListVehicles(t *testing.T) {
	type fields struct {
		Client stubPubSub
	}
	type args struct {
		ctx   context.Context
		tmCfg tm.Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]ha.Device
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Client: stubPubSub{
					subscribeHandler: func(cb paho.MessageHandler) {
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/1/display_name",
							payload: []byte("test-display-name-1"),
						})
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/1/model",
							payload: []byte("test-model-1"),
						})
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/1/trim_badging",
							payload: []byte("test-trim-badging-1"),
						})
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/1/version",
							payload: []byte("test-version-1"),
						})
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/2/display_name",
							payload: []byte("test-display-name-2"),
						})
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/2/trim_badging",
							payload: []byte("test-trim-badging-2"),
						})
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/2/model",
							payload: []byte("test-model-2"),
						})
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/2/version",
							payload: []byte("test-version-2"),
						})

						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/3/trim_badging",
							payload: []byte("test-trim-badging-3"),
						})
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/3/model",
							payload: []byte("test-model-3"),
						})
						cb(nil, &stubMessage{
							topic:   "test-prefix/cars/3/version",
							payload: []byte("test-version-3"),
						})
					},
					subscribeTokens: []paho.Token{&stubToken{}},
				},
			},
			args: args{
				ctx:   context.Background(),
				tmCfg: tm.Config{Prefix: "test-prefix"},
			},
			want: map[string]ha.Device{
				"1": {
					Identifiers:     []string{"test-prefix/cars/1"},
					Manufacturer:    "Tesla",
					Model:           "Model test-model-1 test-trim-badging-1",
					Name:            "test-display-name-1",
					SoftwareVersion: "test-version-1",
					SuggestedArea:   "Garage",
				},
				"2": {
					Identifiers:     []string{"test-prefix/cars/2"},
					Manufacturer:    "Tesla",
					Model:           "Model test-model-2 test-trim-badging-2",
					Name:            "test-display-name-2",
					SoftwareVersion: "test-version-2",
					SuggestedArea:   "Garage",
				},
				"3": {
					Identifiers:     []string{"test-prefix/cars/3"},
					Manufacturer:    "Tesla",
					Model:           "Model test-model-3 test-trim-badging-3",
					Name:            "Tesla",
					SoftwareVersion: "test-version-3",
					SuggestedArea:   "Garage",
				},
			},
		},
		{
			name: "error",
			fields: fields{
				Client: stubPubSub{
					subscribeTokens: []paho.Token{
						&stubToken{err: fmt.Errorf("subscribe error")},
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				tmCfg: tm.Config{Prefix: "test-prefix"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MQTT{
				Client: &tt.fields.Client,
			}
			got, err := m.ListVehicles(tt.args.ctx, tt.args.tmCfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("MQTT.ListVehicles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MQTT.ListVehicles() = %v, want %v", got, tt.want)
			}
		})
	}
}
