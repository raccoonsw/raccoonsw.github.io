package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"log"
	"restApiProject/database"
	"restApiProject/handlers"
	"restApiProject/models"
)

type Specification struct {
	Port       string `required:"true"`
	DBUser     string `required:"true"`
	DBPassword string `required:"true"`
	DBHost     string `required:"true"`
	DBName     string `required:"true"`
	DBPort     string `required:"true"`
}

func setupRouter(sqlDB *database.DBModel) *gin.Engine {
	router := gin.Default()
	router.GET("/api/health", handlers.HealthCheck())
	router.POST("/api/items", handlers.CreateItem(sqlDB))
	router.PUT("/api/items/:id", handlers.UpdateItemById(sqlDB))
	router.PUT("/api/items/sku/:sku", handlers.UpdateItemBySku(sqlDB))
	router.DELETE("/api/items/:id", handlers.DeleteItemById(sqlDB))
	router.DELETE("/api/items/sku/:sku", handlers.DeleteItemBySku(sqlDB))
	router.GET("/api/items/:id", handlers.GetItemById(sqlDB))
	router.GET("/api/items/sku/:sku", handlers.GetItemBySku(sqlDB))
	router.GET("/api/items", handlers.GetAllItems(sqlDB))
	router.POST("/api/graphql", handlers.GraphqlGetItem(sqlDB))
	return router
}

func main() {
	var s Specification
	err := envconfig.Process("restapiproject", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	config := models.Config{DBUser: s.DBUser, DBPassword: s.DBPassword, DBHost: s.DBHost, DBPort: s.DBPort, DBName: s.DBName}
	sqlDB := database.DBModel{DB: database.Connect(config)}
	defer sqlDB.Close()

	router := setupRouter(&sqlDB)
	_ = router.Run(":" + s.Port)
}
