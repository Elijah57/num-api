package main

import (
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"num-api/controllers"
)


func main() {
	
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/classify-number", controllers.ClassifyNumber)

	router.Run(":5080")
}
