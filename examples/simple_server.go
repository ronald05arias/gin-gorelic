package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronald05arias/gin-gorelic"
)

func helloServer(context *gin.Context) {
	context.String(http.StatusOK, "%s", "Hello World")
}

func main() {
	fmt.Println("Starting up")
	router := gin.Default()

	handler, err := gorelic.InitNewrelicAgent("YOUR_NEW_RELIC_LICENSE_KEY", "YOUR_NEW_RELIC_APP_NAME", true, nil)
	if err == nil {
		router.Use(handler)
	} else {
		fmt.Printf("Something went wrong initialising NewRelic %v\n", err)
	}

	router.GET("/", helloServer)
	router.Run(":8080")
}
