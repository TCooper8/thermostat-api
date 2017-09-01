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

var validOperatingModes = map[string] bool {
  "cool": true,
  "heat": true,
  "off": true,
}

var validFanModes = map[string] bool {
  "auto": true,
  "on": true,
}

func validateThermostat(v model.Thermostat) error {
  _, ok := validOperatingModes[v.OperatingMode]
  if !ok {
    return errors.New(fmt.Sprintf("'%v' is not a valid operating mode."))
  }

  _, ok = validFanModes[v.FanMode]
  if !ok {
    return errors.New(fmt.Sprintf("'%s' is not a valid fan mode."))
  }

  return nil
}

func (db *Db) Init() {
  db.Thermostats = make([]model.Thermostat, 0)
}

func (db *Db) FindAll(limit, offset int) []interface{} {
  log.Printf("Querying for thermostats with limit=%v&offset=%v", limit, offset)

  items := make([]interface{}, 0, limit)
  for i := offset; i < limit && i < len(db.Thermostats); i++ {
    items = append(items, db.Thermostats[i])
  }

  return items
}

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

// This function will return true or false if the valud is found.
// The error result signifies a client error -- This needs to be updated to reflect possible database errors.
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
