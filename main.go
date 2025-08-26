package main

import (
	"crud-gin/models"
	"crud-gin/controllers/usercontroller"



	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/api/users", usercontroller.Index)
	router.GET("/api/users/:id", usercontroller.Show)
	router.POST("/api/users", usercontroller.Create)
	router.PUT("/api/users/:id", usercontroller.Update)
	router.DELETE("/api/users/:id", usercontroller.Delete)

	router.Run()
}