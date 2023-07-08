// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package mqtt

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/nebhale/teslamate-discovery/ha"
	"github.com/nebhale/teslamate-discovery/tm"
)

func (m *MQTT) ListVehicles(ctx context.Context, tmCfg tm.Config) (map[string]ha.Device, error) {
	fmt.Println("Listing Vehicles")

	topic := fmt.Sprintf("%s/#", tmCfg.Prefix)

	in, err := m.Subscribe(ctx, topic)
	if err != nil {
		return nil, err
	}
	defer func() { _ = m.Unsubscribe(ctx, topic) }()

	vehicles := make(map[string]ha.Device)
	r := regexp.MustCompile(fmt.Sprintf(`%s/cars/([\d]+)/([\w]+)`, tmCfg.Prefix))

	for {
		select {
		case <-ctx.Done():
			return nil, nil

		case msg := <-in:
			s := r.FindStringSubmatch(msg.Topic())
			if s == nil {
				continue
			}

			switch s[2] {
			case "display_name":
				dev := vehicles[s[1]]
				dev.Name = string(msg.Payload())
				vehicles[s[1]] = dev
			case "model":
				dev := vehicles[s[1]]
				dev.Model = fmt.Sprintf("Model %s%s", string(msg.Payload()), dev.Model)
				vehicles[s[1]] = dev
			case "trim_badging":
				dev := vehicles[s[1]]
				dev.Model = fmt.Sprintf("%s %s", dev.Model, string(msg.Payload()))
				vehicles[s[1]] = dev
			case "version":
				dev := vehicles[s[1]]
				dev.SoftwareVersion = string(msg.Payload())
				vehicles[s[1]] = dev
			}

		case <-time.After(250 * time.Millisecond):
			for id, dev := range vehicles {
				if dev.Name == "" {
					dev.Name = "Tesla"
				}
				dev.Identifiers = []string{fmt.Sprintf("%s/cars/%s", tmCfg.Prefix, id)}
				dev.Manufacturer = "Tesla"
				dev.SuggestedArea = "Garage"
				vehicles[id] = dev
			}

			return vehicles, nil
		}
	}
}
