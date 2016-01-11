package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronald05arias/gin-gorelic"
)

type M struct {
	name        string
	units       string
	directionUp bool
	counter     float64
}

func (_m *M) GetName() string  { return "CustomCounters/" + _m.name }
func (_m *M) GetUnits() string { return _m.units }
func (_m *M) GetValue() (float64, error) {
	if _m.directionUp {
		_m.counter--
		_m.directionUp = false
	} else {
		_m.counter++
		_m.directionUp = true
	}
	return _m.counter, nil
}

func helloServer(context *gin.Context) {
	context.String(http.StatusOK, "%s", "Hello World")
}

func main() {
	fmt.Println("Starting up")
	router := gin.Default()

	m1 := &M{name: "custom_metric_1", units: "secs", counter: 1, directionUp: false}
	m2 := &M{name: "custom_metric_2", units: "secs", counter: 2, directionUp: true}

	handler, err := gorelic.InitNewrelicAgent("YOUR_NEW_RELIC_LICENSE_KEY", "YOUR_NEW_RELIC_APP_NAME", true, []gorelic.Metric{m1, m2})
	if err == nil {
		router.Use(handler)
	} else {
		fmt.Printf("Something went wrong initialising NewRelic %v\n", err)
	}

	router.GET("/", helloServer)
	router.Run(":8080")
}
