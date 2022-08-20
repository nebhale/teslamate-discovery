// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package ha

type Sensor struct {
	Device                 Device            `json:"device,omitempty"`
	DeviceClass            SensorDeviceClass `json:"device_class,omitempty"`
	Icon                   string            `json:"icon,omitempty"`
	JSONAttributesTemplate string            `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string            `json:"json_attributes_topic,omitempty"`
	Name                   string            `json:"name,omitempty"`
	StateClass             StateClass        `json:"state_class,omitempty"`
	StateTopic             string            `json:"state_topic"`
	UniqueId               string            `json:"unique_id,omitempty"`
	UnitOfMeasurement      string            `json:"unit_of_measurement,omitempty"`
	ValueTemplate          string            `json:"value_template,omitempty"`
}

type SensorDeviceClass string

const (
	ApparentPower                         SensorDeviceClass = "apparent_power"
	AirQualityIndex                       SensorDeviceClass = "aqi"
	BatteryCharge                         SensorDeviceClass = "battery"
	CarbonDioxideConcentration            SensorDeviceClass = "carbon_dioxide"
	CarbonMonoxideConcentration           SensorDeviceClass = "carbon_monoxide"
	Current                               SensorDeviceClass = "current"
	Date                                  SensorDeviceClass = "date"
	Duration                              SensorDeviceClass = "duration"
	Energy                                SensorDeviceClass = "energy"
	Frequency                             SensorDeviceClass = "frequency"
	GasVolume                             SensorDeviceClass = "gas"
	Humidity                              SensorDeviceClass = "humidity"
	Illuminance                           SensorDeviceClass = "illuminance"
	Monetary                              SensorDeviceClass = "monetary"
	NitrogenDioxideConcentration          SensorDeviceClass = "nitrogen_dioxide"
	NitrogenMonoxideConcentration         SensorDeviceClass = "nitrogen_monoxide"
	NitrousOxideConcentration             SensorDeviceClass = "nitrous_oxide"
	OzoneConcentration                    SensorDeviceClass = "ozone"
	ParticulateMatter1Concentration       SensorDeviceClass = "pm1"
	ParticulateMatter10Concentration      SensorDeviceClass = "pm10"
	ParticulateMatter25Concentration      SensorDeviceClass = "pm25"
	PowerFactor                           SensorDeviceClass = "power_factor"
	Power                                 SensorDeviceClass = "power"
	Pressure                              SensorDeviceClass = "pressure"
	ReactivePower                         SensorDeviceClass = "reactive_power"
	SignalStrength                        SensorDeviceClass = "signal_strength"
	SulphurDioxideConcentration           SensorDeviceClass = "sulphur_dioxide"
	Temperature                           SensorDeviceClass = "temperature"
	Timestamp                             SensorDeviceClass = "timestamp"
	VolatileOrganicCompoundsConcentration SensorDeviceClass = "volatile_organic_compounds"
	Voltage                               SensorDeviceClass = "voltage"
)

type StateClass string

const (
	Measurement     StateClass = "measurement"
	Total           StateClass = "total"
	TotalIncreasing StateClass = "total_increasing"
)
