# Teslamate Discovery
If you're a fan of the very excellent [Teslamate][tm] and use it with [Home Assistant][ha], then this project is for you!

The Teslamate + Home Assistant [integration documentation][tmha] shows a very long, manual configuration of each of the entities that Teslamate sends messages for.  While this works quite well (I used it for years) it has some shortcomings, the biggest of which is that they aren't all collected under a single [Home Assistant Device][had].  Unfortunately, devices cannot be created using purely manual end-user configuration, they can only be added with [MQTT Discovery][hamd] or a dedicated integration.

This script replaces most, if not all, of the documented Teslamate + Home Assistant integration.  In a single command, it synthesizes and publishes the MQTT Discovery configuration for all direct entities, imperial unit equivalents where appropriate, a couple of useful synthetic `binary_sensors` (e.g. charging, parking break), and a `device_tracker` with precise coordinates that can also determine if the car is home (any geofence with "Home" in it).

[ha]: https://www.home-assistant.io
[had]: https://developers.home-assistant.io/docs/device_registry_index/
[hamd]: https://www.home-assistant.io/docs/mqtt/discovery/
[tm]: https://github.com/adriankumpf/teslamate
[tmha]: https://docs.teslamate.org/docs/integrations/home_assistant

# Pre-requisites
The script assumes that you have a healthy Teslamate installation sending messages to a [Mosquitto MQTT broker][mos] and that Home Assistant can see those messages.  The only other requirement is to have an account that can also send messages to the `/homeassistant/#` topic tree.  This is often done by configuring `/share/mosquitto/accesscontrollist` as in the following example:

```plain
user homeassistant
topic readwrite homeassistant/#
topic read teslamate/#

user teslamate
topic readwrite teslamate/#

user teslamate-discovery
topic readwrite homeassistant/#
```

[mos]: https://github.com/home-assistant/addons/blob/master/mosquitto/DOCS.md

# Usage
The script is relatively self-contained and only depends on `jq` and `mosquitto_pub`. Both utilities are available in a standard Home Assistant install.  The script is commonly copied (or `git clone`'d if you're feeling fancy) into a Home Assistant config directory and executed there.  There are a number of required parameters you'll need to provide, but the script will guide you towards what you're missing.

```plain
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
                              (optional, default: 1883)
  -u, --mqtt-username string  The username to connect to Mosquito MQTT broker with.
                              This user must have at least 'topic write homeassistant/#'
                              access on the broker.
  -P, --mqtt-password string  The password to connect to Mosquito MQTT broker with
  -d, --ha-discovery-prefix   The discovery prefix that Home Assistant is configured
                              to watch. (optional, default: homeassistant)

  --help                      Print this help message
```

## License
Apache License v2.0: see [LICENSE](./LICENSE) for details.
