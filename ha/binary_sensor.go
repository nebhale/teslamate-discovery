// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package ha

type BinarySensor struct {
	Device                 Device                  `json:"device,omitempty"`
	DeviceClass            BinarySensorDeviceClass `json:"device_class,omitempty"`
	Icon                   string                  `json:"icon,omitempty"`
	JSONAttributesTemplate string                  `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string                  `json:"json_attributes_topic,omitempty"`
	Name                   string                  `json:"name,omitempty"`
	PayloadOff             string                  `json:"payload_off,omitempty"`
	PayloadOn              string                  `json:"payload_on,omitempty"`
	StateTopic             string                  `json:"state_topic"`
	UniqueId               string                  `json:"unique_id,omitempty"`
	ValueTemplate          string                  `json:"value_template,omitempty"`
}

type BinarySensorDeviceClass string

const (
	Battery         BinarySensorDeviceClass = "battery"
	BatteryCharging BinarySensorDeviceClass = "battery_charging"
	CarbonMonoxide  BinarySensorDeviceClass = "carbon_monoxide"
	Cold            BinarySensorDeviceClass = "cold"
	Connectivity    BinarySensorDeviceClass = "connectivity"
	Door            BinarySensorDeviceClass = "door"
	Garage          BinarySensorDeviceClass = "garage_door"
	Gas             BinarySensorDeviceClass = "gas"
	Heat            BinarySensorDeviceClass = "heat"
	Light           BinarySensorDeviceClass = "light"
	Lock            BinarySensorDeviceClass = "lock"
	Moisture        BinarySensorDeviceClass = "moisture"
	Motion          BinarySensorDeviceClass = "motion"
	Moving          BinarySensorDeviceClass = "moving"
	Occupancy       BinarySensorDeviceClass = "occupancy"
	Opening         BinarySensorDeviceClass = "opening"
	Plug            BinarySensorDeviceClass = "plug"
	PowerDetected   BinarySensorDeviceClass = "power"
	Presence        BinarySensorDeviceClass = "presence"
	Problem         BinarySensorDeviceClass = "problem"
	Running         BinarySensorDeviceClass = "running"
	Safety          BinarySensorDeviceClass = "safety"
	Smoke           BinarySensorDeviceClass = "smoke"
	Sound           BinarySensorDeviceClass = "sound"
	Tamper          BinarySensorDeviceClass = "tamper"
	Update          BinarySensorDeviceClass = "update"
	Vibration       BinarySensorDeviceClass = "vibration"
	Window          BinarySensorDeviceClass = "window"
)
