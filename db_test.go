package main

import (
	"github.com/satori/go.uuid"
	"github.com/tcooper8/thermostat-api/db"
	"github.com/tcooper8/thermostat-api/model"
	"math/rand"
	"reflect"
	"testing"
)

func TestFindAllEmpty(test *testing.T) {
	db := db.Db{}
	db.Init()
	values := db.FindAll(0, len(db.Thermostats))

	if len(values) != len(db.Thermostats) {
		test.Errorf("FindAll did not return empty array as expected, got %v", values)
	}
}

func TestFindAllN(test *testing.T) {
	db := db.Db{}
	db.Init()

	for i := 0; i < 10; i++ {
		value := model.Thermostat{
			uuid.NewV4(),
			"bob1",
			rand.Float64()*30 + 70,
			"off",
			rand.Float64()*30 + 70,
			rand.Float64()*30 + 70,
			"auto",
		}
		db.Thermostats = append(db.Thermostats, value)
	}

	values := db.FindAll(len(db.Thermostats), 0)
	if len(values) != len(db.Thermostats) {
		test.Errorf("FindAll expected to return %v elements, but got %v.", len(db.Thermostats), len(values))
	}

	for i, other := range db.Thermostats {
		value := values[i]
		if other != value {
			test.Errorf("FindAll expected %v but returned %v", other, value)
		}
	}
}

func TestFindNone(test *testing.T) {
	db := db.Db{}
	db.Init()

	value := db.Find(uuid.NewV4())
	if value != nil {
		test.Errorf("Value was expected to be nil, but got %v", value)
	}
}

func TestFindSome(test *testing.T) {
	db := db.Db{}
	db.Init()
	expected := model.Thermostat{
		uuid.NewV4(),
		"bob1",
		rand.Float64()*30 + 70,
		"off",
		rand.Float64()*30 + 70,
		rand.Float64()*30 + 70,
		"auto",
	}
	db.Thermostats = append(db.Thermostats, expected)

	value := db.Find(expected.Id)
	if value == nil {
		test.Errorf("Value was expected to be %v, but got nil", expected)
	}
	if reflect.DeepEqual(expected, *value) == false {
		test.Errorf("Value was expected to be %v, but got %v", value, expected)
	}
}
