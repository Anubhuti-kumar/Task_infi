package main

import (
	"fmt"

	config "infilon_test/Config"
	handler "infilon_test/Handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	router := gin.Default()

	// GET User
	router.GET("/person/:person_id/info", handler.GetPersonInfo)

	// POST User
	router.POST("/person/create", handler.CreatePerson)

	fmt.Println("Server running on :8080")
	router.Run(":8080")
}
