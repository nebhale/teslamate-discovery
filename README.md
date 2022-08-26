# TeslaMate Discovery
If you're a fan of the very excellent [TeslaMate][tm] and use it with [Home Assistant][ha], then this project is for you!

The TeslaMate + Home Assistant [integration documentation][tmha] shows a very long, manual configuration of each of the entities that TeslaMate sends messages for.  While this works quite well (I used it for years) it has some shortcomings, the biggest of which is that they aren't all collected under a single [Home Assistant Device][had].  Unfortunately, devices cannot be created using purely manual end-user configuration, they can only be added with [MQTT Discovery][hamd] or a dedicated integration.

This application replaces most, if not all, of the documented TeslaMate + Home Assistant integration.  In a single command, it synthesizes and publishes the MQTT Discovery configuration for commonly used entities converted to desired units and a `device_tracker` with precise coordinates that can also determine if the car is home (any geofence with "Home" in it).

[ha]: https://www.home-assistant.io
[had]: https://developers.home-assistant.io/docs/device_registry_index/
[hamd]: https://www.home-assistant.io/docs/mqtt/discovery/
[tm]: https://github.com/adriankumpf/teslamate
[tmha]: https://docs.teslamate.org/docs/integrations/home_assistant

# Pre-requisites
The application assumes that you have a healthy TeslaMate installation sending messages to a [Mosquitto MQTT broker][mos] and that Home Assistant can see those messages.  The only other requirement is to have an account that can read messages from the `/teslamate/#` topic tree and send messages to the `/homeassistant/#` topic tree.  This is often done by configuring `/share/mosquitto/accesscontrollist` as in the following example:

```plain
user homeassistant
topic readwrite homeassistant/#
topic read teslamate/#

user teslamate
topic write teslamate/#

user teslamate-discovery
topic write homeassistant/#
topic read teslamate/#
```

[mos]: https://github.com/home-assistant/addons/blob/master/mosquitto/DOCS.md

# Usage
```plain
Usage:
  teslamate-discovery [flags]

Flags:
      --ha-discovery-prefix string   home assistant discovery message prefix (default "homeassistant")
      --help                         help for teslamate-discovery
  -h, --mqtt-host string             mqtt broker host (default "127.0.0.1")
  -P, --mqtt-password string         mqtt broker password
  -p, --mqtt-port int                mqtt broker port (default 8883)
  -s, --mqtt-scheme string           mqtt broker scheme (default "ssl")
  -u, --mqtt-username string         mqtt broker username
      --tm-prefix string             teslamate message prefix (default "teslamate")
      --units-distance string        distance units ["imperial", "metric"] (default "imperial")
      --units-pressure string        pressure units ["imperial", "metric"] (default "imperial")
  -v, --version                      version for teslamate-discovery
```

## License
Apache License v2.0: see [LICENSE](./LICENSE) for details.
