package model

import (
  "github.com/satori/go.uuid"
)

// Operating mode should be one of "cool" | "heat" | "off"
// FanMode should be one of "auto" | "on"
type Thermostat struct {
  Id uuid.UUID `json:"id"`
  Name string `json:"name"`
  CurrentTemp float64 `json:"temp"`
  OperatingMode string `json:"operatingMode"`
  CoolSetPoint float64 `json:"coolSetPoint"`
  HeatSetPoint float64 `json:"heatSetPoint"`
  FanMode string `json:"fanMode"`
}

type ThermostatPatch struct {
  Name string `json:"name"`
  OperatingMode string `json:"operatingMode"`
  CoolSetPoint float64 `json:"coolSetPoint"`
  HeatSetPoint float64 `json:"heatSetPoint"`
  FanMode string `json:"fanMode"`
}
