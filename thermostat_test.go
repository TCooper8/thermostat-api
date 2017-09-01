package main

import (
  "github.com/satori/go.uuid"
  "math/rand"
  "encoding/json"
  "bytes"
  "fmt"
  "os"
  "net/http/httptest"
  "net/http"
  "testing"
  "github.com/tcooper8/thermostat-api/model"
	"github.com/tcooper8/thermostat-api/db"
	"github.com/tcooper8/thermostat-api/server"
)

var app server.App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
  rr := httptest.NewRecorder()
  app.Router.ServeHTTP(rr, req)

  return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
  if expected != actual {
    t.Errorf("Expected response code %d. Got %d\n", expected, actual)
  }
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

func TestMain(m *testing.M) {
  app = server.App{}
  db := &db.Db{}

  db.Init()
  db.Thermostats = thermostats
  app.Init(db)

  code := m.Run()
  os.Exit(code)
}

func TestListThermostats(test *testing.T) {
  req, _ := http.NewRequest("GET", "/hub/thermostats", nil)
  response := executeRequest(req)
  checkResponseCode(test, http.StatusOK, response.Code)

  var values []model.Thermostat
  json.Unmarshal(response.Body.Bytes(), &values)
  if len(values) != len(thermostats) {
    test.Errorf("Expected values to be %v, got %v", thermostats, values)
  }
  for i := range values {
    if values[i] != thermostats[i] {
      test.Errorf("Expected values to be %v, got %v", thermostats, values)
    }
  }
}

func TestListThermostatsOffset(test *testing.T) {
  for i, _ := range thermostats {
    offset := i
    count := len(thermostats) - offset
    query := fmt.Sprintf("?offset=%v", offset)
    req, _ := http.NewRequest("GET", "/hub/thermostats" + query, nil)
    response := executeRequest(req)
    checkResponseCode(test, http.StatusOK, response.Code)

    var values []model.Thermostat
    json.Unmarshal(response.Body.Bytes(), &values)

    if len(values) != count {
      test.Errorf("Expected values to be length %v, but got %v", count, len(values))
    }
  }
}

func TestListThermostatsLimit(test *testing.T) {
  limit := 1
  query := fmt.Sprintf("?limit=%v", limit)
  req, _ := http.NewRequest("GET", "/hub/thermostats" + query, nil)
  response := executeRequest(req)
  checkResponseCode(test, http.StatusOK, response.Code)

  var values []model.Thermostat
  json.Unmarshal(response.Body.Bytes(), &values)

  if len(values) != limit {
    test.Errorf("Expected values to be limit to %v, but got %v", limit, len(values))
  }
}

func TestGetThermostatNotFound(test *testing.T) {
  id := uuid.NewV4()
  req, _ := http.NewRequest("GET", "/hub/thermostats/" + id.String(), nil)
  response := executeRequest(req)
  checkResponseCode(test, http.StatusNotFound, response.Code)
}

func TestGetThermostat(test *testing.T) {
  for _, thermostat := range thermostats {
    id := thermostat.Id
    req, _ := http.NewRequest("GET", "/hub/thermostats/" + id.String(), nil)
    response := executeRequest(req)
    checkResponseCode(test, http.StatusOK, response.Code)

    var value model.Thermostat
    json.Unmarshal(response.Body.Bytes(), &value)
    if value != thermostat {
      test.Errorf("Expected id %v to get thermostat %v, but got %v", id, thermostat, value)
    }
  }
}

func TestPatchThermostatNotFound(test *testing.T) {
  id := uuid.NewV4()
  patch := model.ThermostatPatch{
    "test",
    "cool",
    0.0,
    0.0,
    "auto",
  }
  patchBytes, _ := json.Marshal(patch)

  req, _ := http.NewRequest("PATCH", "/hub/thermostats/" + id.String(), bytes.NewBuffer(patchBytes))
  response := executeRequest(req)
  checkResponseCode(test, http.StatusNotFound, response.Code)
}

func TestPatchThermostatNoInput(test *testing.T) {
  var payload []byte
  for _, thermostat := range thermostats {
    id := thermostat.Id
    req, _ := http.NewRequest("PATCH", "/hub/thermostats/" + id.String(), bytes.NewBuffer(payload))
    resp := executeRequest(req)
    checkResponseCode(test, http.StatusBadRequest, resp.Code)
  }
}

func TestPatchThermostat(test *testing.T) {
  patch := model.ThermostatPatch{
    "test",
    "cool",
    0.0,
    0.0,
    "auto",
  }
  patchBytes, _ := json.Marshal(patch)

  for _, thermostat := range thermostats {
    id := thermostat.Id
    req, _ := http.NewRequest("PATCH", "/hub/thermostats/" + id.String(), bytes.NewBuffer(patchBytes))
    resp := executeRequest(req)
    checkResponseCode(test, http.StatusNoContent, resp.Code)
  }
}

func TestPatchThermostatBadFanMode(test *testing.T) {
  patch := model.ThermostatPatch{
    "test",
    "cool",
    0.0,
    0.0,
    "bad mode!",
  }
  patchBytes, _ := json.Marshal(patch)

  for _, thermostat := range thermostats {
    id := thermostat.Id
    req, _ := http.NewRequest("PATCH", "/hub/thermostats/" + id.String(), bytes.NewBuffer(patchBytes))
    resp := executeRequest(req)
    checkResponseCode(test, http.StatusBadRequest, resp.Code)
  }
}

func TestPatchThermostatBadOperatingMode(test *testing.T) {
  patch := model.ThermostatPatch{
    "test",
    "bad mode!",
    0.0,
    0.0,
    "auto",
  }
  patchBytes, _ := json.Marshal(patch)

  for _, thermostat := range thermostats {
    id := thermostat.Id
    req, _ := http.NewRequest("PATCH", "/hub/thermostats/" + id.String(), bytes.NewBuffer(patchBytes))
    resp := executeRequest(req)
    checkResponseCode(test, http.StatusBadRequest, resp.Code)
  }
}
