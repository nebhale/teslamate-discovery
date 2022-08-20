// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package mqtt

import (
	"context"
	"encoding/json"
	"fmt"

	paho "github.com/eclipse/paho.mqtt.golang"

	"github.com/nebhale/teslamate-discovery/ha"
)

type MQTT struct {
	Client paho.Client
}

func NewMQTT(ctx context.Context, config Config) (*MQTT, error) {
	uri := BrokerURI(config)
	fmt.Printf("Connecting to %s\n", uri)

	c := paho.NewClient(paho.NewClientOptions().
		AddBroker(uri).
		SetClientID(fmt.Sprintf("teslamate-discovery-%s", RandomString(12))).
		SetOrderMatters(false).
		SetUsername(config.Username).
		SetPassword(config.Password))

	t := c.Connect()
	select {
	case <-ctx.Done():
		return nil, nil
	case <-t.Done():
		if err := t.Error(); err != nil {
			return nil, err
		}
	}

	go func() {
		<-ctx.Done()
		c.Disconnect(500)
	}()

	return &MQTT{Client: c}, nil
}

func (m *MQTT) Publish(ctx context.Context, discoveryPrefix string, v ...interface{}) error {
	for _, v := range v {
		var topic string

		switch v := v.(type) {
		case ha.BinarySensor:
			topic = fmt.Sprintf("%s/binary_sensor/%s/config", discoveryPrefix, v.UniqueId)
		case ha.DeviceTracker:
			topic = fmt.Sprintf("%s/device_tracker/%s/config", discoveryPrefix, v.UniqueId)
		case ha.Sensor:
			topic = fmt.Sprintf("%s/sensor/%s/config", discoveryPrefix, v.UniqueId)
		default:
			return fmt.Errorf("unexpected message type: %T", v)
		}

		fmt.Printf("  %s\n", topic)

		payload, err := json.Marshal(v)
		if err != nil {
			return err
		}

		t := m.Client.Publish(topic, 0, true, payload)
		select {
		case <-ctx.Done():
			return nil
		case <-t.Done():
			if err := t.Error(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *MQTT) Subscribe(ctx context.Context, topic string) (<-chan paho.Message, error) {
	ch := make(chan paho.Message)

	t := m.Client.Subscribe(topic, 0, func(c paho.Client, m paho.Message) {
		ch <- m
	})

	select {
	case <-ctx.Done():
		return nil, nil
	case <-t.Done():
		if err := t.Error(); err != nil {
			return nil, err
		}
	}

	return ch, nil
}

func (m *MQTT) Unsubscribe(ctx context.Context, topic string) error {
	t := m.Client.Unsubscribe(topic)

	select {
	case <-ctx.Done():
		return nil
	case <-t.Done():
		if err := t.Error(); err != nil {
			return err
		}
	}

	return nil
}

func BrokerURI(config Config) string {
	return fmt.Sprintf("%s://%s:%d", config.Scheme, config.Host, config.Port)
}
