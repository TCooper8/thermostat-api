package db

import (
  "github.com/satori/go.uuid"
  "github.com/tcooper8/thermostat-api/model"
  "log"
)

type Db struct {
  Thermostats []model.Thermostat
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

func (db *Db) Patch(id uuid.UUID, patch model.ThermostatPatch) bool {
  for i, v := range db.Thermostats {
    if v.Id == id {
      log.Printf("Updating %v with %v", v, patch)
      db.Thermostats[i] = model.Thermostat{
        v.Id,
        patch.Name,
        v.CurrentTemp,
        patch.OperatingMode,
        patch.CoolSetPoint,
        patch.HeatSetPoint,
        patch.FanMode,
      }
      return true
    }
  }

  return false
}
