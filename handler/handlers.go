package handler

import (
  "fmt"
  "github.com/satori/go.uuid"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/tcooper8/thermostat-api/db"
  "github.com/tcooper8/thermostat-api/model"
  "strconv"
  "log"
  "net/http"
)

type Handler struct {
  db *db.Db
}

func writeJsonResponse(req *http.Request, resp http.ResponseWriter, bytes []byte) {
  log.Printf("Response %v with %v:%v", req.URL, 200)
  resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
  resp.Write(bytes)
}

func httpError(req *http.Request, resp http.ResponseWriter, err string, code int) {
  log.Printf("Response %v with %v:%v", req.URL, err, code)
  http.Error(resp, err, code)
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
    httpError(req, resp, err.Error(), http.StatusBadRequest)
  }
  offset, err := getOffset(0, req)
  if err != nil {
    httpError(req, resp, err.Error(), http.StatusBadRequest)
  }

  thermostats := handler.db.FindAll(limit, offset)
  bytes, err := json.Marshal(thermostats)
  if err != nil {
    httpError(req, resp, err.Error(), http.StatusInternalServerError)
  } else {
    resp.Header().Set("x-total-count", strconv.Itoa(len(thermostats)))
    resp.Header().Set("link", fmt.Sprintf("/hub/thermostats?offset=%v&limit=%v", offset + limit, limit))
    writeJsonResponse(req, resp, bytes)
  }
}

func (handler *Handler) GetThermostat(resp http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  id, err := uuid.FromString(vars["uuid"])
  if err != nil {
    httpError(req, resp, err.Error(), http.StatusBadRequest)
    return
  }
  thermostat := handler.db.Find(id)
  if thermostat == nil {
    httpError(req, resp, "Thermostat not found.", http.StatusNotFound)
    return
  }

  bytes, err := json.Marshal(thermostat)
  if err != nil {
    httpError(req, resp, err.Error(), http.StatusInternalServerError)
  }

  writeJsonResponse(req, resp, bytes)
}

func (handler *Handler) PatchThermostat(resp http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  id, err := uuid.FromString(vars["uuid"])
  if err != nil {
    httpError(req, resp, err.Error(), http.StatusBadRequest)
    return
  }

  decoder := json.NewDecoder(req.Body)
  defer req.Body.Close()
  if req.Body == nil {
    httpError(req, resp, "Expected body to be non-null.", http.StatusBadRequest)
  }

  patch := model.ThermostatPatch{}
  err = decoder.Decode(&patch)
  if err != nil {
    httpError(req, resp, err.Error(), http.StatusBadRequest)
  }

  updated, badRequest := handler.db.Patch(id, patch)
  if badRequest != nil {
    httpError(req, resp, badRequest.Error(), http.StatusBadRequest)
  } else if !updated {
    httpError(req, resp, "Thermostat not found.", http.StatusNotFound)
  } else {
    httpError(req, resp, "", http.StatusNoContent)
  }
}
