// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package ha

type DeviceTracker struct {
	Device                 Device                  `json:"device,omitempty"`
	Icon                   string                  `json:"icon,omitempty"`
	JSONAttributesTemplate string                  `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string                  `json:"json_attributes_topic,omitempty"`
	Name                   string                  `json:"name,omitempty"`
	PayloadHome            string                  `json:"payload_home,omitempty"`
	PayloadNotHome         string                  `json:"payload_not_home,omitempty"`
	SourceType             DeviceTrackerSourceType `json:"source_type,omitempty"`
	StateTopic             string                  `json:"state_topic"`
	UniqueId               string                  `json:"unique_id,omitempty"`
	ValueTemplate          string                  `json:"value_template,omitempty"`
}

type DeviceTrackerSourceType string

const (
	Bluetooth   DeviceTrackerSourceType = "bluetooth"
	BluetoothLE DeviceTrackerSourceType = "bluetooth_le"
	GPS         DeviceTrackerSourceType = "gps"
	Router      DeviceTrackerSourceType = "router"
)
