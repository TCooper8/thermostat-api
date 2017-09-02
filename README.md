# Just a dummy api for Golang!

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

## REST API

**List Thermostats**
----
  Lists a range of thermostats, specified with `limit` and `offset`.

* **URL**
  /hub/thermostats
* **METHOD:**
  `GET`
* **URL Params**
  **Optional:**
  `offset={int}`
  `limit={int}`

* **Success Response**
  * **Code:** 200 <br />
    **Content:** `[model.Thermostat]`

* **Error Response**

* **Sample Call**
  `GET /hub/thermostats?offset=0&limit=5`

**Get Thermostat**
----
  Will attempt to retrive a specific thermostat.

* **URL**
  `/hub/thermostats/<thermostat.id: uuid>```
* **METHOD:**
  GET
* **Success Response**
  * **Code:** 200 <br />
    **Content:** `model.Thermostat`
* **Error Response**
  * **Code:** 404 <br />


**Patch Thermostat**
----
  Will attempt to update a given record by id, with the values of the input structure -- Does not currently protect against undefined or default values if fields are ommited. Non-write fields are ignored.

* **URL**
  `/hub/thermostats/<thermostat.id: uuid>`
* **METHOD:**
  PATCH
* **Data Params**
  ```json
  {
    "name": "bob1",
    "operatingMode": "off" | "cool" | "heat",
    "coolSetPoint": 75,
    "heatSetPoint": 72,
    "fanMode": "auto" | "on"
  }
  ```
* **Success Response**
  * **Code:** 204 <br />
* **Error Response**
  * **Code:** 404 <br />
