package main

import (
	"github.com/satori/go.uuid"
	"github.com/tcooper8/thermostat-api/db"
	"github.com/tcooper8/thermostat-api/model"
	"github.com/tcooper8/thermostat-api/server"
	"log"
	"math/rand"
	"os"
)

func main() {
	// This is all dummy data for the in-memory data store.
	// This should be replaced with a real database implementation.
	var thermostats = []model.Thermostat{
		model.Thermostat{
			uuid.NewV4(),
			"bob1",
			rand.Float64()*30 + 70,
			"off",
			rand.Float64()*30 + 70,
			rand.Float64()*30 + 70,
			"auto",
		},
		model.Thermostat{
			uuid.NewV4(),
			"bob2",
			rand.Float64()*30 + 70,
			"heat",
			rand.Float64()*30 + 70,
			rand.Float64()*30 + 70,
			"auto",
		},
		model.Thermostat{
			uuid.NewV4(),
			"bob3",
			rand.Float64()*30 + 70,
			"off",
			rand.Float64()*30 + 70,
			rand.Float64()*30 + 70,
			"auto",
		},
		model.Thermostat{
			uuid.NewV4(),
			"bob4",
			rand.Float64()*30 + 70,
			"cool",
			rand.Float64()*30 + 70,
			rand.Float64()*30 + 70,
			"auto",
		},
		model.Thermostat{
			uuid.NewV4(),
			"bob5",
			rand.Float64()*30 + 70,
			"off",
			rand.Float64()*30 + 70,
			rand.Float64()*30 + 70,
			"auto",
		},
	}

	app := server.App{}
	db := &db.Db{}

	db.Init()
	db.Thermostats = thermostats
	app.Init(db)

	host := os.Getenv("HOST")
	if host == "" {
		host = ":8080"
	}

	log.Printf("Started on %v", host)
	app.Run(host)
}
