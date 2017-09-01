package main

import (
  "github.com/satori/go.uuid"
  "math/rand"
  "github.com/tcooper8/thermostat-api/model"
	"github.com/tcooper8/thermostat-api/db"
	"github.com/tcooper8/thermostat-api/server"
)

func main() {
  // This is all dummy data for the in-memory data store.
  // This should be replaced with a real database implementation.
  var thermostats = []model.Thermostat {
    model.Thermostat {
      uuid.NewV4(),
      "bob1",
      rand.Float64() * 20 + 70,
      "off",
      rand.Float64() * 20 + 70,
      rand.Float64() * 20 + 70,
      "auto",
    },
    model.Thermostat {
      uuid.NewV4(),
      "bob2",
      rand.Float64() * 20 + 70,
      "heat",
      rand.Float64() * 20 + 70,
      rand.Float64() * 20 + 70,
      "auto",
    },
    model.Thermostat {
      uuid.NewV4(),
      "bob3",
      rand.Float64() * 20 + 70,
      "off",
      rand.Float64() * 20 + 70,
      rand.Float64() * 20 + 70,
      "auto",
    },
    model.Thermostat {
      uuid.NewV4(),
      "bob4",
      rand.Float64() * 20 + 70,
      "cool",
      rand.Float64() * 20 + 70,
      rand.Float64() * 20 + 70,
      "auto",
    },
    model.Thermostat {
      uuid.NewV4(),
      "bob5",
      rand.Float64() * 20 + 70,
      "off",
      rand.Float64() * 20 + 70,
      rand.Float64() * 20 + 70,
      "auto",
    },
  }

  app := server.App{}
  db := &db.Db{}

  db.Init()
  db.Thermostats = thermostats
  app.Init(db)

  app.Run(":8080")
}
