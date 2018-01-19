package api

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	mid "api/middleware"
	"HighFive"
)

// Starts API
func StartAPI(five HighFive.HighFive) {
	ConfigRuntime()
	StartGin(five)
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}


// Sets up API files and routing functions
func StartGin(five HighFive.HighFive) {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "resources/static")
	router.GET("/", mid.Index)
	router.GET("/basic", mid.Basic)

	router.GET("/getFive", mid.GetFive(five))
	router.POST("/search", mid.SendResults(five))
	router.GET("/receive", mid.ReceiveAjax)

	router.Run("localhost:8080")
}