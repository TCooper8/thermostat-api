package handler

import (
  "github.com/satori/go.uuid"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/tcooper8/thermostat-api/db"
  "github.com/tcooper8/thermostat-api/model"
  "strconv"
  "net/http"
)

type Handler struct {
  db *db.Db
}

func writeJsonResponse(resp http.ResponseWriter, bytes []byte) {
  resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
  resp.Write(bytes)
}

func (handler *Handler) Init(db *db.Db) {
  handler.db = db
}

func getOffset(defaultValue int, req *http.Request) (int, error) {
  query := req.URL.Query()
  offsets, ok := query["offset"]
  offset := defaultValue
  var err error

  if ok {
    if len(offsets) > 0 {
      offsetStr := offsets[0]
      offset, err = strconv.Atoi(offsetStr)
      if err != nil {
        return 0, err
      }
    }
  }

  return offset, nil
}

func getLimit(maxValue, defaultValue int, req *http.Request) (int, error) {
  query := req.URL.Query()
  limits, ok := query["limit"]
  limit := defaultValue
  var err error

  if ok {
    if len(limits) > 0 {
      limitStr := limits[0]
      limit, err = strconv.Atoi(limitStr)
      if err != nil {
        return 0, err
      }

      if limit > maxValue {
        limit = maxValue
      }
    }
  }

  return limit, nil
}

func (handler *Handler) GetThermostats(resp http.ResponseWriter, req *http.Request) {
  limit, err := getLimit(100, 10, req)
  if err != nil {
    http.Error(resp, err.Error(), http.StatusBadRequest)
  }
  offset, err := getOffset(0, req)
  if err != nil {
    http.Error(resp, err.Error(), http.StatusBadRequest)
  }

  thermostats := handler.db.FindAll(limit, offset)
  bytes, err := json.Marshal(thermostats)
  if err != nil {
    http.Error(resp, err.Error(), http.StatusInternalServerError)
  }

  writeJsonResponse(resp, bytes)
}

func (handler *Handler) GetThermostat(resp http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  id, err := uuid.FromString(vars["uuid"])
  if err != nil {
    http.Error(resp, err.Error(), http.StatusBadRequest)
    return
  }
  thermostat := handler.db.Find(id)
  if thermostat == nil {
    http.Error(resp, "Thermostat not found.", http.StatusNotFound)
    return
  }

  bytes, err := json.Marshal(thermostat)
  if err != nil {
    http.Error(resp, err.Error(), http.StatusInternalServerError)
  }

  writeJsonResponse(resp, bytes)
}

func (handler *Handler) PatchThermostat(resp http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  id, err := uuid.FromString(vars["uuid"])
  if err != nil {
    http.Error(resp, err.Error(), http.StatusBadRequest)
    return
  }

  decoder := json.NewDecoder(req.Body)
  defer req.Body.Close()
  if req.Body == nil {
    http.Error(resp, "Expected body to be non-null.", http.StatusBadRequest)
  }

  patch := model.ThermostatPatch{}
  err = decoder.Decode(&patch)
  if err != nil {
    http.Error(resp, err.Error(), http.StatusBadRequest)
  }

  badRequest, err := handler.db.Patch(id, patch)
  if badRequest != nil {
    http.Error(resp, badRequest.Error(), http.StatusBadRequest)
  } else if err != nil {
    http.Error(resp, err.Error(), http.StatusBadRequest)
  } else {
    http.Error(resp, "", http.StatusNoContent)
  }
}
