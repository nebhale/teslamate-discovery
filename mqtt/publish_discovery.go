// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package mqtt

import (
	"context"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/nebhale/teslamate-discovery/ha"
	"github.com/nebhale/teslamate-discovery/units"
)

func (m *MQTT) PublishDiscovery(ctx context.Context, id string, device ha.Device, haCfg ha.Config,
	unitsCfg units.Config) error {

	fmt.Printf("Configuring %s\n", device.Name)

	v := []interface{}{

		// Charge
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Energy,
			Name:              Name(device, "Energy Added"),
			StateTopic:        StateTopic(device, "/charge_energy_added"),
			UniqueId:          UniqueId(device, "/charge_energy_added"),
			UnitOfMeasurement: "kWh",
			ValueTemplate:     units.RoundingValueTemplate,
		},
		ha.Sensor{
			Device:            device,
			Icon:              "mdi:battery-charging-90",
			Name:              Name(device, "Limit"),
			StateClass:        ha.Measurement,
			StateTopic:        StateTopic(device, "/charge_limit_soc"),
			UniqueId:          UniqueId(device, "/limit"),
			UnitOfMeasurement: "%",
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Current,
			Name:              Name(device, "Charger Current"),
			StateTopic:        StateTopic(device, "/charger_actual_current"),
			UniqueId:          UniqueId(device, "/charger_current"),
			UnitOfMeasurement: "A",
		},
		ha.BinarySensor{
			Device:        device,
			DeviceClass:   ha.BatteryCharging,
			Name:          Name(device, "Charger"),
			StateTopic:    StateTopic(device, "/state"),
			UniqueId:      UniqueId(device, "/charging"),
			ValueTemplate: `{{ "ON" if value == "charging" else "OFF" }}`,
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Plug,
			Name:        Name(device, "Charger"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/plugged_in"),
			UniqueId:    UniqueId(device, "/plug"),
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Power,
			Name:              Name(device, "Charger Power"),
			StateTopic:        StateTopic(device, "/charger_power"),
			UniqueId:          UniqueId(device, "/charger_power"),
			UnitOfMeasurement: "kW",
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Voltage,
			Name:              Name(device, "Charger Voltage"),
			StateTopic:        StateTopic(device, "/charger_voltage"),
			UniqueId:          UniqueId(device, "/charger_voltage"),
			UnitOfMeasurement: "V",
		},
		ha.Sensor{
			Device:      device,
			DeviceClass: ha.Timestamp,
			Name:        Name(device, "Scheduled Start Time"),
			StateTopic:  StateTopic(device, "/scheduled_charging_start_time"),
			UniqueId:    UniqueId(device, "/start_time"),
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Duration,
			Icon:              "mdi:timer",
			Name:              Name(device, "Time to Charged"),
			StateTopic:        StateTopic(device, "/time_to_full_charge"),
			UniqueId:          UniqueId(device, "/time_to_charged"),
			UnitOfMeasurement: "h",
		},

		// Climate
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Temperature,
			Name:              Name(device, "Inside Temp"),
			StateClass:        ha.Measurement,
			StateTopic:        StateTopic(device, "/inside_temp"),
			UniqueId:          UniqueId(device, "/inside_temp"),
			UnitOfMeasurement: "°C",
			ValueTemplate:     units.RoundingValueTemplate,
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Running,
			Icon:        "mdi:fan",
			Name:        Name(device, "Climate"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/is_climate_on"),
			UniqueId:    UniqueId(device, "/climate"),
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Running,
			Icon:        "mdi:fan",
			Name:        Name(device, "Preconditioning"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/is_preconditioning"),
			UniqueId:    UniqueId(device, "/preconditioning"),
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Temperature,
			Name:              Name(device, "Outside Temp"),
			StateClass:        ha.Measurement,
			StateTopic:        StateTopic(device, "/outside_temp"),
			UniqueId:          UniqueId(device, "/outside_temp"),
			UnitOfMeasurement: "°C",
			ValueTemplate:     units.RoundingValueTemplate,
		},

		// Location
		ha.Sensor{
			Device:            device,
			Icon:              "mdi:image-filter-hdr",
			Name:              Name(device, "Elevation"),
			StateTopic:        StateTopic(device, "/elevation"),
			UniqueId:          UniqueId(device, "/elevation"),
			UnitOfMeasurement: unitsCfg.Distance.DistanceShortUnits(),
			ValueTemplate:     unitsCfg.Distance.DistanceShortValueTemplate(),
		},
		ha.Sensor{
			Device:     device,
			Icon:       "mdi:earth",
			Name:       Name(device, "Geofence"),
			StateTopic: StateTopic(device, "/geofence"),
			UniqueId:   UniqueId(device, "/geofence"),
		},
		ha.Sensor{
			Device:            device,
			Icon:              "mdi:compass",
			Name:              Name(device, "Heading"),
			StateTopic:        StateTopic(device, "/heading"),
			UniqueId:          UniqueId(device, "/heading"),
			UnitOfMeasurement: "°",
		},
		ha.Sensor{
			Device:            device,
			Icon:              "mdi:crosshairs-gps",
			Name:              Name(device, "Latitude"),
			StateTopic:        StateTopic(device, "/latitude"),
			UniqueId:          UniqueId(device, "/latitude"),
			UnitOfMeasurement: "°",
		},
		ha.DeviceTracker{
			Device:                 device,
			Icon:                   "mdi:car",
			JSONAttributesTemplate: fmt.Sprintf(`{{ { "latitude": value | float(0), "longitude": states("sensor.%s_longitude") | float(0) } | to_json }}`, strcase.ToSnake(device.Name)),
			JSONAttributesTopic:    StateTopic(device, "/latitude"),
			Name:                   device.Name,
			SourceType:             "gps",
			StateTopic:             StateTopic(device, "/latitude"),
			UniqueId:               UniqueId(device, "/location"),
			ValueTemplate:          fmt.Sprintf(`{{ "home" if "home" in (states("sensor.%s_geofence") | lower) else "not_home" }}`, strcase.ToSnake(device.Name)),
		},
		ha.Sensor{
			Device:            device,
			Icon:              "mdi:crosshairs-gps",
			Name:              Name(device, "Longitude"),
			StateTopic:        StateTopic(device, "/longitude"),
			UniqueId:          UniqueId(device, "/longitude"),
			UnitOfMeasurement: "°",
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Power,
			Name:              Name(device, "Power"),
			StateTopic:        StateTopic(device, "/power"),
			UniqueId:          UniqueId(device, "/power"),
			UnitOfMeasurement: "kW",
		},
		ha.Sensor{
			Device:            device,
			Icon:              "mdi:speedometer",
			Name:              Name(device, "Speed"),
			StateTopic:        StateTopic(device, "/speed"),
			UniqueId:          UniqueId(device, "/speed"),
			UnitOfMeasurement: unitsCfg.Distance.SpeedUnits(),
			ValueTemplate:     unitsCfg.Distance.SpeedValueTemplate(),
		},

		// State
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.BatteryCharge,
			Name:              Name(device, "Battery"),
			StateClass:        ha.Measurement,
			StateTopic:        StateTopic(device, "/battery_level"),
			UniqueId:          UniqueId(device, "/battery"),
			UnitOfMeasurement: "%",
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Door,
			Icon:        "mdi:ev-plug-tesla",
			Name:        Name(device, "Charge Port"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/charge_port_door_open"),
			UniqueId:    UniqueId(device, "/charge_port"),
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Door,
			Icon:        "mdi:car-door",
			Name:        Name(device, "Doors"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/doors_open"),
			UniqueId:    UniqueId(device, "/doors"),
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Door,
			Icon:        "mdi:car",
			Name:        Name(device, "Frunk"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/frunk_open"),
			UniqueId:    UniqueId(device, "/frunk"),
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Problem,
			Icon:        "mdi:heart-pulse",
			Name:        Name(device, "Health"),
			PayloadOff:  "true",
			PayloadOn:   "false",
			StateTopic:  StateTopic(device, "/healthy"),
			UniqueId:    UniqueId(device, "/health"),
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Lock,
			Name:        Name(device, "Locked"),
			PayloadOff:  "true",
			PayloadOn:   "false",
			StateTopic:  StateTopic(device, "/locked"),
			UniqueId:    UniqueId(device, "/locked"),
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Occupancy,
			Icon:        "mdi:account",
			Name:        Name(device, "Occupied"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/is_user_present"),
			UniqueId:    UniqueId(device, "/occupied"),
		},
		ha.Sensor{
			Device:            device,
			Icon:              "mdi:counter",
			Name:              Name(device, "Odometer"),
			StateClass:        ha.TotalIncreasing,
			StateTopic:        StateTopic(device, "/odometer"),
			UniqueId:          UniqueId(device, "/odometer"),
			UnitOfMeasurement: unitsCfg.Distance.DistanceLongUnits(),
			ValueTemplate:     unitsCfg.Distance.DistanceLongValueTemplate(),
		},
		ha.Sensor{
			Device:            device,
			Icon:              "mdi:map-marker-distance",
			Name:              Name(device, "Range"),
			StateClass:        ha.Measurement,
			StateTopic:        StateTopic(device, "/rated_battery_range_km"),
			UniqueId:          UniqueId(device, "/range"),
			UnitOfMeasurement: unitsCfg.Distance.DistanceLongUnits(),
			ValueTemplate:     unitsCfg.Distance.DistanceLongValueTemplate(),
		},
		ha.BinarySensor{
			Device:     device,
			Icon:       "mdi:cctv",
			Name:       Name(device, "Sentry Mode"),
			PayloadOff: "false",
			PayloadOn:  "true",
			StateTopic: StateTopic(device, "/sentry_mode"),
			UniqueId:   UniqueId(device, "/sentry_mode"),
		},
		ha.Sensor{
			Device:     device,
			Icon:       "mdi:car-connected",
			Name:       Name(device, "State"),
			StateTopic: StateTopic(device, "/state"),
			UniqueId:   UniqueId(device, "/state"),
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Pressure,
			Icon:              "mdi:car-tire-alert",
			Name:              Name(device, "Tire Pressure (Front Left)"),
			StateClass:        ha.Measurement,
			StateTopic:        StateTopic(device, "/tpms_pressure_fl"),
			UniqueId:          UniqueId(device, "/tire_pressure_front_left"),
			UnitOfMeasurement: unitsCfg.Pressure.PressureUnits(),
			ValueTemplate:     unitsCfg.Pressure.PressureValueTemplate(),
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Pressure,
			Icon:              "mdi:car-tire-alert",
			Name:              Name(device, "Tire Pressure (Front Right)"),
			StateClass:        ha.Measurement,
			StateTopic:        StateTopic(device, "/tpms_pressure_fr"),
			UniqueId:          UniqueId(device, "/tire_pressure_front_right"),
			UnitOfMeasurement: unitsCfg.Pressure.PressureUnits(),
			ValueTemplate:     unitsCfg.Pressure.PressureValueTemplate(),
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Pressure,
			Icon:              "mdi:car-tire-alert",
			Name:              Name(device, "Tire Pressure (Rear Left)"),
			StateClass:        ha.Measurement,
			StateTopic:        StateTopic(device, "/tpms_pressure_rl"),
			UniqueId:          UniqueId(device, "/tire_pressure_rear_left"),
			UnitOfMeasurement: unitsCfg.Pressure.PressureUnits(),
			ValueTemplate:     unitsCfg.Pressure.PressureValueTemplate(),
		},
		ha.Sensor{
			Device:            device,
			DeviceClass:       ha.Pressure,
			Icon:              "mdi:car-tire-alert",
			Name:              Name(device, "Tire Pressure (Rear Right)"),
			StateClass:        ha.Measurement,
			StateTopic:        StateTopic(device, "/tpms_pressure_rr"),
			UniqueId:          UniqueId(device, "/tire_pressure_rear_right"),
			UnitOfMeasurement: unitsCfg.Pressure.PressureUnits(),
			ValueTemplate:     unitsCfg.Pressure.PressureValueTemplate(),
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Door,
			Icon:        "mdi:car",
			Name:        Name(device, "Trunk"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/trunk_open"),
			UniqueId:    UniqueId(device, "/trunk"),
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Update,
			Name:        Name(device, "Update"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/update_available"),
			UniqueId:    UniqueId(device, "/update"),
		},
		ha.BinarySensor{
			Device:      device,
			DeviceClass: ha.Window,
			Icon:        "mdi:car-door",
			Name:        Name(device, "Windows"),
			PayloadOff:  "false",
			PayloadOn:   "true",
			StateTopic:  StateTopic(device, "/windows_open"),
			UniqueId:    UniqueId(device, "/windows"),
		},
		ha.Sensor{
			Device:     device,
			Icon:       "mdi:numeric",
			Name:       Name(device, "Version"),
			StateTopic: StateTopic(device, "/version"),
			UniqueId:   UniqueId(device, "/version"),
		},
	}

	return m.Publish(ctx, haCfg.DiscoveryPrefix, v...)
}

func Name(device ha.Device, suffix string) string {
	return fmt.Sprintf("%s %s", device.Name, suffix)
}

func StateTopic(device ha.Device, suffix string) string {
	return fmt.Sprintf("%s%s", device.Identifiers[0], suffix)
}

func UniqueId(device ha.Device, suffix string) string {
	return fmt.Sprintf("%s%s", strings.ReplaceAll(device.Identifiers[0], "/", "_"), suffix)
}
