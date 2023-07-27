package main

import (
	"assignment/src/configs"
	"assignment/src/db"
	"assignment/src/models"
	"assignment/src/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.InitEnvConfig()
	log.Println("Initiated Config", configs.EnvConfigs)
	db.SetupDB()
	database := db.GetDB()
	database.AutoMigrate(&models.Product{})
	database.AutoMigrate(&models.User{})

	log.Println("DB initaiated", database)
	router := gin.Default()

	routes.InitializeRouter(router)
	log.Println("Router initiated")

	router.Run(configs.EnvConfigs.LocalServerPort)

}
