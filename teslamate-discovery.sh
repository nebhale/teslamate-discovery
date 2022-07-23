#!/usr/bin/env bash
# Copyright 2022 Ben Hale
# SPDX-License-Identifier: Apache-2.0
# shellcheck disable=SC2016

set -euo pipefail

binary_sensor() {
  local SENSOR
  SENSOR=$(snake_case "$1")
  local UNIQUE_SENSOR
  UNIQUE_SENSOR=$(unique_sensor "$@")

  jq \
    --argjson ADDITIONAL "$(jq --null-input --arg CAR_NAME "${CAR_NAME}" "$2")" \
    --argjson DEVICE "$(device)" \
    --arg CAR_ID "${CAR_ID}" \
    --arg CAR_NAME "${CAR_NAME}" \
    --arg NAME "$1" \
    --arg SENSOR "${SENSOR}" \
    --arg UNIQUE_SENSOR "${UNIQUE_SENSOR}" \
    --null-input \
    --compact-output \
    '{
      device: $DEVICE,
      name: "\($CAR_NAME) \($NAME)",
      payload_off: "false",
      payload_on: "true",
      state_topic: "teslamate/cars/\($CAR_ID)/\($SENSOR)",
      unique_id: "teslamate/cars/\($CAR_ID)/\($UNIQUE_SENSOR)"
    } + $ADDITIONAL' | publish binary_sensor "${UNIQUE_SENSOR}"
}

device() {
  jq \
    --null-input \
    --arg CAR_ID "${CAR_ID}" \
    --arg CAR_MODEL "${CAR_MODEL}" \
    --arg CAR_NAME "${CAR_NAME}" \
    --arg CAR_SW_VERSION "${CAR_SW_VERSION}" \
    '{
      identifiers: "teslamate/cars/\($CAR_ID)",
      manufacturer: "Tesla",
      model: $CAR_MODEL,
      name: $CAR_NAME,
      suggested_area: "Garage",
      sw_version: $CAR_SW_VERSION
    }'
}

device_tracker() {
  local SENSOR
  SENSOR=$(snake_case "$1")
  local UNIQUE_SENSOR
  UNIQUE_SENSOR=$(unique_sensor "$@")

  jq \
    --argjson ADDITIONAL "$(jq --null-input --arg CAR_NAME "${CAR_NAME}" "$2")" \
    --argjson DEVICE "$(device)" \
    --arg CAR_ID "${CAR_ID}" \
    --arg CAR_NAME "${CAR_NAME}" \
    --arg NAME "$1" \
    --arg SC_CAR_NAME "$(snake_case "${CAR_NAME}")" \
    --arg SENSOR "${SENSOR}" \
    --arg UNIQUE_SENSOR "${UNIQUE_SENSOR}" \
    --null-input \
    --compact-output \
    '{
      device: $DEVICE,
      name: "\($CAR_NAME) \($NAME)",
      source_type: "gps",
      json_attributes_topic: "teslamate/cars/\($CAR_ID)/\($SENSOR)",
      json_attributes_template: "{{ { \"latitude\": value | float(0), \"longitude\": states(\"sensor.\($SC_CAR_NAME)_longitude\") | float(0), \"gps_accuracy\": 1 } | to_json }}",
      state_topic: "teslamate/cars/\($CAR_ID)/\($SENSOR)",
      unique_id: "teslamate/cars/\($CAR_ID)/\($UNIQUE_SENSOR)",
      value_template: "{{ \"home\" if \"home\" in (states(\"sensor.\($SC_CAR_NAME)_geofence\") | lower) else \"not_home\" }}"
    } + $ADDITIONAL' | publish device_tracker "${UNIQUE_SENSOR}"
}

print_usage() {
  cat <<EOF
Usage:
  teslamate-discovery.sh [flags]

Flags:
  -i, --car-id number         The teslamate id for the car, e.g. 1
  -m, --car-model string      The model of the car, e.g. Model S 85D
  -n, --car-name string       The name of the car, e.g. My Beautiful Blue Baby
  -v, --car-sw-version string The software version of the car, e.g. 2022.16.3
  -h, --mqtt-host string      The host of the Mosquito MQTT broker to connect to.
                              (optional, default: core-mosquitto)
  -p, --mqtt-port number      The port of the Mosquito MQTT broker to connect to.
                              (optional, default: 8883)
  -u, --mqtt-username string  The username to connect to Mosquito MQTT broker with.
                              This user must have at least 'topic write homeassistant/#'
                              access on the broker.
  -P, --mqtt-password string  The password to connect to Mosquito MQTT broker with
  -d, --ha-discovery-prefix   The discovery prefix that Home Assistant is configured
                              to watch. (optional, default: homeassistant)

  --help                      Print this help message
EOF
}

publish() {
  echo "Publishing to ${HA_DISCOVERY_PREFIX}/$1/teslamate_cars_${CAR_ID}/$2/config"

  mosquitto_pub \
    --host "${MQTT_HOST}" \
    --port "${MQTT_PORT}" \
    --username "${MQTT_USERNAME}" \
    --pw "${MQTT_PASSWORD}" \
    --topic "${HA_DISCOVERY_PREFIX}/$1/teslamate_cars_${CAR_ID}/$2/config" \
    --retain \
    --stdin-file
}

sensor() {
  local SENSOR
  SENSOR=$(snake_case "$1")
  local UNIQUE_SENSOR
  UNIQUE_SENSOR=$(unique_sensor "$@")

  jq \
    --argjson ADDITIONAL "$(jq --null-input --arg CAR_NAME "${CAR_NAME}" "$2")" \
    --argjson DEVICE "$(device)" \
    --arg CAR_ID "${CAR_ID}" \
    --arg CAR_NAME "${CAR_NAME}" \
    --arg NAME "$1" \
    --arg SENSOR "${SENSOR}" \
    --arg UNIQUE_SENSOR "${UNIQUE_SENSOR}" \
    --null-input \
    --compact-output \
    '{
      device: $DEVICE,
      name: "\($CAR_NAME) \($NAME)",
      state_topic: "teslamate/cars/\($CAR_ID)/\($SENSOR)",
      unique_id: "teslamate/cars/\($CAR_ID)/\($UNIQUE_SENSOR)"
    } + $ADDITIONAL' | publish sensor "${UNIQUE_SENSOR}"
}

snake_case() {
  tr ' ' '_' <<<"$1"  | tr '[:upper:]' '[:lower:]'
}

unique_sensor() {
  if [[ -z ${3+x} ]]; then
    snake_case "$1"
  else
    snake_case "$3"
  fi
}

while [[ $# -gt 0 ]]; do
  case $1 in
    -i | --car-id)
      CAR_ID="$2"
      shift # past argument
      shift # past value
      ;;
    -m | --car-model)
      CAR_MODEL="$2"
      shift # past argument
      shift # past value
      ;;
    -n | --car-name)
      CAR_NAME="$2"
      shift # past argument
      shift # past value
      ;;
    -v | --car-sw-version)
      CAR_SW_VERSION="$2"
      shift # past argument
      shift # past value
      ;;
    -h | --mqtt-host)
      MQTT_HOST="$2"
      shift # past argument
      shift # past value
      ;;
    -p | --mqtt-port)
      MQTT_PORT="$2"
      shift # past argument
      shift # past value
      ;;
    -u | --mqtt-username)
      MQTT_USERNAME="$2"
      shift # past argument
      shift # past value
      ;;
    -P | --mqtt-password)
      MQTT_PASSWORD="$2"
      shift # past argument
      shift # past value
      ;;
    -d | --ha-discovery-prefix)
      HA_DISCOVERY_PREFIX="$2"
      shift # past argument
      shift # past value
      ;;
    --help)
      print_usage
      exit 0
      ;;
    -*)
      printf -- "Unknown option %s\n\n" "$1"
      print_usage
      exit 1
      ;;
  esac
done

if [[ -z ${CAR_ID+x} ]]; then
  printf -- "-i, --car-id must be specified\n\n"
  print_usage
  exit 1
fi

if [[ -z ${CAR_MODEL+x} ]]; then
  printf -- "-m, --car-model must be specified\n\n"
  print_usage
  exit 1
fi

if [[ -z ${CAR_NAME+x} ]]; then
  printf -- "-n, --car-name must be specified\n\n"
  print_usage
  exit 1
fi

if [[ -z ${CAR_SW_VERSION+x} ]]; then
  printf -- "-v, --car-sw-version must be specified\n\n"
  print_usage
  exit 1
fi

if [[ -z ${MQTT_HOST+x} ]]; then
  MQTT_HOST="core-mosquitto"
fi

if [[ -z ${MQTT_PORT+x} ]]; then
  MQTT_PORT="8883"
fi

if [[ -z ${MQTT_USERNAME+x} ]]; then
  printf -- "-u, --mqtt-username must be specified\n\n"
  print_usage
  exit 1
fi

if [[ -z ${MQTT_PASSWORD+x} ]]; then
  printf -- "-P, --mqtt-password must be specified\n\n"
  print_usage
  exit 1
fi

if [[ -z ${HA_DISCOVERY_PREFIX+x} ]]; then
  HA_DISCOVERY_PREFIX="homeassistant"
fi

sensor        "Display Name"                  '{ icon: "mdi:car-electric" }'
sensor        "State"                         '{ icon: "mdi:car-connected" }'
sensor        "Since"                         '{ device_class: "timestamp" }'
binary_sensor "Healthy"                       '{ icon: "mdi:heart-pulse" }'
sensor        "Version"                       '{ icon: "mdi:numeric" }'
binary_sensor "Update Available"              '{ device_class: "update" }'
sensor        "Update Version"                '{ icon: "mdi:numeric" }'

sensor        "Model"                         '{ icon: "mdi:car-electric" }'
sensor        "Trim Badging"                  '{ icon: "mdi:shield-star" }'
sensor        "Exterior Color"                '{ icon: "mdi:palette" }'
sensor        "Wheel Type"                    '{ icon: "mdi:tire" }'
sensor        "Spoiler Type"                  '{ icon: "mdi:car-sports" }'

sensor        "Geofence"                      '{ icon: "mdi:earth" }'

sensor        "Latitude"                      '{ icon: "mdi:crosshairs-gps", unit_of_measurement: "°" }'
sensor        "Longitude"                     '{ icon: "mdi:crosshairs-gps", unit_of_measurement: "°" }'
sensor        "Shift State"                   '{ icon: "mdi:car-shift-pattern" }'
sensor        "Power"                         '{ device_class: "power", unit_of_measurement: "kW" }'
sensor        "Speed"                         '{ icon: "mdi:speedometer", unit_of_measurement:"km/h" }'
sensor        "Speed"                         '{ icon: "mdi:speedometer", unit_of_measurement:"mph", value_template: "{{ (value | float(0) / 1.609344) | round(2) }}" }' "speed_mph"
sensor        "Heading"                       '{ icon: "mdi:compass", unit_of_measurement: "°" }'
sensor        "Elevation"                     '{ icon: "mdi:image-filter-hdr", unit_of_measurement: "m" }'
sensor        "Elevation"                     '{ icon: "mdi:image-filter-hdr", unit_of_measurement: "ft", value_template: "{{ (value | float(0) * 3.280839) | round(2) }}" }' "elevation_ft"

binary_sensor "Locked"                        '{ device_class: "lock", value_template: "{{ \"false\" if value == \"true\" else \"true\" }}" }'
binary_sensor "Sentry Mode"                   '{ icon: "mdi:cctv" }'
binary_sensor "Windows Open"                  '{ name: "\($CAR_NAME) Windows", device_class: "window", icon: "mdi:car-door" }'
binary_sensor "Doors Open"                    '{ name: "\($CAR_NAME) Doors", device_class: "door", icon: "mdi:car-door" }'
binary_sensor "Trunk Open"                    '{ name: "\($CAR_NAME) Trunk", device_class: "opening", icon: "mdi:car-back" }'
binary_sensor "Frunk Open"                    '{ name: "\($CAR_NAME) Frunk", device_class: "opening", icon: "mdi:car" }'
binary_sensor "Is User Present"               '{ name: "\($CAR_NAME) User Presence", device_class: "occupancy", icon: "mdi:account" }'

binary_sensor "Is Climate On"                 '{ name: "\($CAR_NAME) Climate", device_class: "running", icon: "mdi:fan" }'
sensor        "Inside Temp"                   '{ name: "\($CAR_NAME) Inside Temperature", device_class: "temperature", unit_of_measurement: "°C" }'
sensor        "Outside Temp"                  '{ name: "\($CAR_NAME) Outside Temperature", device_class: "temperature", unit_of_measurement: "°C" }'
binary_sensor "Is Preconditioning"            '{ name: "\($CAR_NAME) Preconditioning", device_class: "running", icon: "mdi:fan" }'

sensor        "Odometer"                      '{ icon: "mdi:counter", unit_of_measurement: "km" }'
sensor        "Odometer"                      '{ icon: "mdi:counter", unit_of_measurement: "mi", value_template: "{{ (value | float(0) / 1.609344) | round(2) }}" }' "odometer_mi"
sensor        "Est Battery Range KM"          '{ name: "\($CAR_NAME) Estimated Battery Range", icon: "mdi:gauge", unit_of_measurement: "km" }'
sensor        "Est Battery Range KM"          '{ name: "\($CAR_NAME) Estimated Battery Range", icon: "mdi:gauge", unit_of_measurement: "mi", value_template: "{{ (value | float(0) / 1.609344) | round(2) }}" }' "est_battery_range_mi"
sensor        "Rated Battery Range KM"        '{ name: "\($CAR_NAME) Rated Battery Range", icon: "mdi:gauge", unit_of_measurement: "km" }'
sensor        "Rated Battery Range KM"        '{ name: "\($CAR_NAME) Rated Battery Range", icon: "mdi:gauge", unit_of_measurement: "mi", value_template: "{{ (value | float(0) / 1.609344) | round(2) }}" }' "rated_battery_range_mi"
sensor        "Ideal Battery Range KM"        '{ name: "\($CAR_NAME) Ideal Battery Range", icon: "mdi:gauge", unit_of_measurement: "km" }'
sensor        "Ideal Battery Range KM"        '{ name: "\($CAR_NAME) Ideal Battery Range", icon: "mdi:gauge", unit_of_measurement: "mi", value_template: "{{ (value | float(0) / 1.609344) | round(2) }}" }' "ideal_battery_range_mi"

sensor        "Battery Level"                 '{ device_class: "battery", unit_of_measurement: "%" }'
sensor        "Usable Battery Level"          '{ device_class: "battery", unit_of_measurement: "%" }'
binary_sensor "Plugged In"                    '{ name: "\($CAR_NAME) Plug", device_class: "plug", icon: "mdi:ev-station" }'
sensor        "Charge Energy Added"           '{ device_class: "energy", icon: "mdi:battery-charging-50", unit_of_measurement: "kWh" }'
sensor        "Charge Limit SOC"              '{ icon: "mdi:battery-charging-90", unit_of_measurement: "%" }'
binary_sensor "Charge Port Door Open"         '{ device_class: "door", icon: "mdi:ev-plug-tesla" }'
sensor        "Charger Actual Current"        '{ device_class: "current", unit_of_measurement: "A" }'
sensor        "Charger Phases"                '{ icon: "mdi:sine-wave" }'
sensor        "Charger Power"                 '{ device_class: "power", unit_of_measurement: "kW" }'
sensor        "Charger Voltage"               '{ device_class: "voltage", unit_of_measurement: "V" }'
sensor        "Charge Current Request"        '{ device_class: "current", unit_of_measurement: "A" }'
sensor        "Charge Current Request Max"    '{ name: "\($CAR_NAME) Charge Current Request Maximum", device_class: "current", unit_of_measurement: "A" }'
sensor        "Scheduled Charging Start Time" '{ device_class: "timestamp" }'
sensor        "Time to Full Charge"           '{ device_class: "duration", unit_of_measurement: "hours" }'

sensor        "TPMS Pressure FL"              '{ name: "\($CAR_NAME) Tire Pressure (Front Left)", device_class: "pressure", unit_of_measurement: "bar" }'
sensor        "TPMS Pressure FL"              '{ name: "\($CAR_NAME) Tire Pressure (Front Left)", device_class: "pressure", unit_of_measurement: "psi", value_template: "{{ (value | float(0) * 14.503773) | round(2) }}" }' "tpms_pressure_fl_psi"
sensor        "TPMS Pressure FR"              '{ name: "\($CAR_NAME) Tire Pressure (Front Right)", device_class: "pressure", unit_of_measurement: "bar" }'
sensor        "TPMS Pressure FR"              '{ name: "\($CAR_NAME) Tire Pressure (Front Right)", device_class: "pressure", unit_of_measurement: "psi", value_template: "{{ (value | float(0) * 14.503773) | round(2) }}" }' "tpms_pressure_fr_psi"
sensor        "TPMS Pressure RL"              '{ name: "\($CAR_NAME) Tire Pressure (Rear Left)", device_class: "pressure", unit_of_measurement: "bar" }'
sensor        "TPMS Pressure RL"              '{ name: "\($CAR_NAME) Tire Pressure (Rear Left)", device_class: "pressure", unit_of_measurement: "psi", value_template: "{{ (value | float(0) * 14.503773) | round(2) }}" }' "tpms_pressure_rl_psi"
sensor        "TPMS Pressure RR"              '{ name: "\($CAR_NAME) Tire Pressure (Rear Right)", device_class: "pressure", unit_of_measurement: "bar" }'
sensor        "TPMS Pressure RR"              '{ name: "\($CAR_NAME) Tire Pressure (Rear Right)", device_class: "pressure", unit_of_measurement: "psi", value_template: "{{ (value | float(0) * 14.503773) | round(2) }}" }' "tpms_pressure_rr_psi"

binary_sensor "State"                         '{ name: "\($CAR_NAME) Charging", device_class: "battery_charging", value_template: "{{ \"true\" if value == \"charging\" else \"false\" }}" }' "charging"
binary_sensor "Shift State"                   '{ name: "\($CAR_NAME) Parking Brake", icon: "mdi:car-brake-parking", value_template: "{{ \"true\" if value == \"P\" else \"false\" }}" }' "parking_brake"

device_tracker "Latitude"                     '{ name: "\($CAR_NAME)", icon: "mdi:car" }' "location"
