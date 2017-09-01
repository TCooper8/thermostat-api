# Just a dummy api for Golang!

## REST API

## Authorization

Currently, there is none.

### GET /hub/thermostats?offset=<offset>&limit=<limit>

Will list a range of thermostats within offset and limit.

### GET /hub/thermostats/<thermostat.id: uuid>

Will attempt to retrive a specific value.

### PATCH /hug/thermostats/<thermostat.id: uuid> with <body>

```json
{
  "name": "bob1",
  "operatingMode": "off",
  "coolSetPoint": 75,
  "heatSetPoint": 72,
  "fanMode": "off"
}
```

Will attempt to update a given record by id, with the values of the input structure -- Does not currently protect against undefined or default values if fields are ommited.
