package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"log"
	"restApiProject/controllers"
	"restApiProject/models"
)

type Specification struct {
	Port       string `required:"true"`
	DBUser     string `required:"true"`
	DBPassword string `required:"true"`
	DBHost     string `required:"true"`
	DBName     string `required:"true"`
}

func setupRouter(sqlDB models.DBModel) *gin.Engine {
	router := gin.Default()
	env := &controllers.Env{DBModel: sqlDB}
	router.GET("/health", env.HealthCheck)
	router.POST("/item", env.CreateItem)
	router.PUT("/item/:id", env.UpdateItemById)
	router.PUT("/item/sku/:sku", env.UpdateItemBySku)
	router.DELETE("/item/:id", env.DeleteItemById)
	router.DELETE("/item/sku/:sku", env.DeleteItemBySku)
	router.GET("/item/:id", env.GetItemById)
	router.GET("/item/sku/:sku", env.GetItemBySku)
	router.GET("/items", env.GetAllItems)
	return router
}

func main() {
	var s Specification
	err := envconfig.Process("restapiproject", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	config := models.Config{DBUser: s.DBUser, DBPassword: s.DBPassword, DBHost: s.DBHost, DBName: s.DBName}
	sqlDB := models.DBModel{DB: models.Connect(config)}
	defer sqlDB.Close()

	router := setupRouter(sqlDB)
	_ = router.Run(":" + s.Port)
}
