package main

import (
	"github.com/gin-gonic/gin"
	"restApiProject/controllers"
	"restApiProject/models"
)

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
	config := models.Config{User: "root", Password: "g7y48UPH", DBName: "restApi"}
	sqlDB := models.DBModel{DB: models.Connect(config)}
	defer sqlDB.Close()

	router := setupRouter(sqlDB)
	_ = router.Run()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
