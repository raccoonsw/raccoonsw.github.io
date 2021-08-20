package database

import (
	"restApiProject/models"
)

type OrdersInterface interface {
	CreateOrder(models.Order) (models.Order, error)
}

func (sqlDB *DBModel) CreateOrder(order models.Order) (models.Order, error) {
	result := sqlDB.DB.Create(&order)
	return order, result.Error
}
