package db

import (
  "github.com/satori/go.uuid"
  "github.com/tcooper8/thermostat-api/model"
  "log"
  "fmt"
  "errors"
)

type Db struct {
  Thermostats []model.Thermostat
}

// This may need to be replaced with a database operation in the future, for sanitizing user input for Thermostat.OperatingMode.
var validOperatingModes = map[string] bool {
  "cool": true,
  "heat": true,
  "off": true,
}

// This may need to be replaced with a database operation in the future, for sanitizing user input for Thermostat.FanMode.
var validFanModes = map[string] bool {
  "auto": true,
  "on": true,
}

// Used to validate user input for a Thermostat.
// This may need to be updated with input validation for temperature values as well. Obviously we don't want someone to set CoolSetPoint to an extremely low value.
func validateThermostat(v model.Thermostat) error {
  _, ok := validOperatingModes[v.OperatingMode]
  if !ok {
    return errors.New(fmt.Sprintf("'%v' is not a valid operating mode.", v.OperatingMode))
  }

  _, ok = validFanModes[v.FanMode]
  if !ok {
    return errors.New(fmt.Sprintf("'%s' is not a valid fan mode.", v.FanMode))
  }

  if v.HeatSetPoint < 30.0 || v.HeatSetPoint > 100.0 {
    return errors.New("'HeatSetPoint' is not within the range of [30.0, 100.0]")
  }

  if v.CoolSetPoint < 30.0 || v.CoolSetPoint > 100.0 {
    return errors.New("'CoolSetPoint' is not within the range of [30.0, 100.0]")
  }

  return nil
}

func (db *Db) Init() {
  db.Thermostats = make([]model.Thermostat, 0)
}

// Queries for a range of thermostats, by limit and offset.
func (db *Db) FindAll(limit, offset int) []interface{} {
  log.Printf("Querying for thermostats with limit=%v&offset=%v", limit, offset)

  items := make([]interface{}, 0, limit)
  for i := offset; i < (limit + offset) && i < len(db.Thermostats); i++ {
    items = append(items, db.Thermostats[i])
  }

  return items
}

// Finds a single Thermostat by id.
func (db *Db) Find(id uuid.UUID) *model.Thermostat {
  log.Printf("Querying for thermostat %v", id)

  for _, v := range db.Thermostats {
    if v.Id == id {
      thermostat := &model.Thermostat{
        v.Id,
        v.Name,
        v.CurrentTemp,
        v.OperatingMode,
        v.CoolSetPoint,
        v.HeatSetPoint,
        v.FanMode,
      }
      return thermostat
    }
  }

  return nil
}

// The bool will signify if the value has been updated.
// The first error will signify any bad request.
func (db *Db) Patch(id uuid.UUID, patch model.ThermostatPatch) (bool, error) {
  thermo := model.Thermostat {
    uuid.Nil,
    patch.Name,
    0.0,
    patch.OperatingMode,
    patch.CoolSetPoint,
    patch.HeatSetPoint,
    patch.FanMode,
  }

  // This will check for invalid user input.
  err := validateThermostat(thermo)
  if err != nil {
    return false, err
  }

  for i, v := range db.Thermostats {
    if v.Id == id {
      log.Printf("Updating %v with %v", v, patch)
      thermo.Id = v.Id
      thermo.CurrentTemp = v.CurrentTemp
      db.Thermostats[i] = thermo
      return true, nil
    }
  }

  return false, nil
}
