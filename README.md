# Just a dummy api for Golang!

## REST API

## SETUP

```bash
go build
go test
go run main.go
```

+ env
  - HOST : Example, ":8080"

## Authorization

Currently, there is none.

** Get Thermostats **
----
  Will list a range of thermostats within offset and limit.

* **URL**
``` GET /hub/thermostats?offset=<offset>&limit=<limit> ```

** Get Thermostat **
----
  Will attempt to retrive a specific value.

* **URL**
``` GET /hub/thermostats/<thermostat.id: uuid> ```

** Patch Thermostat **
----
  Will attempt to update a given record by id, with the values of the input structure -- Does not currently protect against undefined or default values if fields are ommited. Non-write fields are ignored.

* **URL**
``` PATCH /hub/thermostats/<thermostat.id: uuid> with <body> ```

```json
{
  "name": "bob1",
  "operatingMode": "off" | "cool" | "heat",
  "coolSetPoint": 75,
  "heatSetPoint": 72,
  "fanMode": "auto" | "on"
}
```
