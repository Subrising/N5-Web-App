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

	// Sets up server files and paths
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/js", "./web/js")

	router.GET("/", mid.Index)
	router.GET("/races/:id", mid.ShowRace)
	router.GET("/receive", mid.NextFive(five))

	router.POST("/raceDetails", mid.GetRace(five))

	router.Run("localhost:8080")
}