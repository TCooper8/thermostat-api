package server

import (
  "github.com/satori/go.uuid"
  "math/rand"
	"github.com/gorilla/mux"
	"github.com/tcooper8/thermostat-api/handler"
  "github.com/tcooper8/thermostat-api/model"
	"github.com/tcooper8/thermostat-api/db"
	"log"
	"net/http"
)

type App struct {
  Router *mux.Router
  Db *db.Db
  Handler *handler.Handler
}

func (app *App) Init (db *db.Db) {
	router := mux.NewRouter().StrictSlash(true)
  handler := &handler.Handler{}
  handler.Init(db)

	sub := router.PathPrefix("/hub").Subrouter()
	sub.Methods("GET").Path("/thermostats").HandlerFunc(handler.GetThermostats)
  sub.Methods("GET").Path("/thermostats/{uuid}").HandlerFunc(handler.GetThermostat)
  sub.Methods("PATCH").Path("/thermostats/{uuid}").HandlerFunc(handler.PatchThermostat)

  app.Router = router
  app.Db = db
  app.Handler = handler
}

func (app *App) Run (endPoint string) {
	log.Fatal(http.ListenAndServe(endPoint, app.Router))
}

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
