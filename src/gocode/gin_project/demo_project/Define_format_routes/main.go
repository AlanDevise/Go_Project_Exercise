package main

import (
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

func main() {
	router := gin.Default()
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	router.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	router.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// Listen and Server in http://0.0.0.0:8080
	router.Run()

}
