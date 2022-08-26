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
)

func TestMQTT_Publish(t *testing.T) {
	type fields struct {
		Client stubPubSub
	}
	type args struct {
		ctx             context.Context
		discoveryPrefix string
		v               []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "known",
			fields: fields{
				Client: stubPubSub{
					publishTokens: []paho.Token{&stubToken{}},
				},
			},
			args: args{
				ctx:             context.Background(),
				discoveryPrefix: "test-discovery-prefix",
				v: []interface{}{
					ha.BinarySensor{
						UniqueId: "test-unique-id",
					},
					ha.DeviceTracker{
						UniqueId: "test-unique-id",
					},
					ha.Sensor{
						UniqueId: "test-unique-id",
					},
				},
			},
			want: []string{
				"test-discovery-prefix/binary_sensor/test-unique-id/config",
				"test-discovery-prefix/device_tracker/test-unique-id/config",
				"test-discovery-prefix/sensor/test-unique-id/config",
			},
		},
		{
			name: "unknown",
			fields: fields{
				Client: stubPubSub{
					publishTokens: []paho.Token{&stubToken{}},
				},
			},
			args: args{
				ctx:             context.Background(),
				discoveryPrefix: "test-discovery-prefix",
				v:               []interface{}{""},
			},
			wantErr: true,
		},
		{
			name: "error",
			fields: fields{
				Client: stubPubSub{
					publishTokens: []paho.Token{
						&stubToken{err: fmt.Errorf("publish error")},
					},
				},
			},
			args: args{
				ctx:             context.Background(),
				discoveryPrefix: "test-discovery-prefix",
				v: []interface{}{
					ha.BinarySensor{
						UniqueId: "test-unique-id",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MQTT{
				Client: &tt.fields.Client,
			}
			if err := m.Publish(tt.args.ctx, tt.args.discoveryPrefix, tt.args.v...); (err != nil) != tt.wantErr {
				t.Errorf("MQTT.Publish() error = %v, wantErr %v", err, tt.wantErr)
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

func TestMQTT_Subscribe(t *testing.T) {
	type fields struct {
		Client stubPubSub
	}
	type args struct {
		ctx   context.Context
		topic string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    paho.Message
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Client: stubPubSub{
					subscribeHandler: func(callback paho.MessageHandler) {
						callback(nil, &stubMessage{messageId: 1})
					},
					subscribeTokens: []paho.Token{&stubToken{}},
				},
			},
			args: args{
				ctx:   context.Background(),
				topic: "test-topic",
			},
			want: &stubMessage{messageId: 1},
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
				topic: "test-topic",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MQTT{
				Client: &tt.fields.Client,
			}
			ch, err := m.Subscribe(tt.args.ctx, tt.args.topic)
			if (err != nil) != tt.wantErr {
				t.Errorf("MQTT.Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			got := <-ch
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MQTT.Subscribe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMQTT_Unsubscribe(t *testing.T) {
	type fields struct {
		Client stubPubSub
	}
	type args struct {
		ctx   context.Context
		topic string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Client: stubPubSub{
					unsubscribeTokens: []paho.Token{&stubToken{}},
				},
			},
			args: args{
				ctx:   context.Background(),
				topic: "test-topic",
			},
			want: []string{"test-topic"},
		},
		{
			name: "error",
			fields: fields{
				Client: stubPubSub{
					unsubscribeTokens: []paho.Token{
						&stubToken{err: fmt.Errorf("unsubscribe error")},
					},
				},
			},
			args: args{
				ctx:   context.Background(),
				topic: "test-topic",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MQTT{
				Client: &tt.fields.Client,
			}
			if err := m.Unsubscribe(tt.args.ctx, tt.args.topic); (err != nil) != tt.wantErr {
				t.Errorf("MQTT.Unsubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			args := tt.fields.Client.unsubscribeArgs[0]
			if !reflect.DeepEqual(args.topics, tt.want) {
				t.Errorf("MQTT.PublishDiscovery() topics = %v, want %v", args.topics, tt.want)
				return
			}
		})
	}
}

func TestBrokerURI(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				config: Config{
					Scheme: "test-scheme",
					Host:   "test-host",
					Port:   4242,
				},
			},
			want: "test-scheme://test-host:4242",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BrokerURI(tt.args.config); got != tt.want {
				t.Errorf("BrokerURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
