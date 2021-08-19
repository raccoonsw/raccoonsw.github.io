package database

import (
	"gorm.io/gorm"
	"restApiProject/models"
)

type ItemsInterface interface {
	CreateItem(models.Item) (models.Item, error)
	UpdateItemById(models.Item) error
	UpdateItemBySku(models.Item) error
	DeleteItemById(int) error
	DeleteItemBySku(string) error
	GetItemById(int) (models.Item, error)
	GetItemBySku(string) (models.Item, error)
	GetAllItems(models.Paging) ([]models.Item, bool, error)
}

type DBModel struct {
	DB *gorm.DB
}

func (sqlDB *DBModel) CreateItem(item models.Item) (models.Item, error) {
	result := sqlDB.DB.Create(&item)
	return item, result.Error
}

func (sqlDB *DBModel) UpdateItemById(item models.Item) error {
	result := sqlDB.DB.Save(&item)
	return result.Error
}

func (sqlDB *DBModel) UpdateItemBySku(item models.Item) error {
	result := sqlDB.DB.Model(&models.Item{}).Where("sku = ?", item.Sku).Updates(item)
	return result.Error
}

func (sqlDB *DBModel) DeleteItemById(id int) error {
	result := sqlDB.DB.Delete(&models.Item{}, id)
	return result.Error
}

func (sqlDB *DBModel) DeleteItemBySku(sku string) error {
	result := sqlDB.DB.Where("sku = ?", sku).Delete(&models.Item{})
	return result.Error
}

func (sqlDB *DBModel) GetItemById(id int) (models.Item, error) {
	var newItem models.Item
	result := sqlDB.DB.First(&newItem, id)
	return newItem, result.Error
}

func (sqlDB *DBModel) GetItemBySku(sku string) (models.Item, error) {
	var newItem models.Item
	result := sqlDB.DB.First(&newItem, "sku = ?", sku)
	return newItem, result.Error
}

func (sqlDB *DBModel) GetAllItems(paging models.Paging) ([]models.Item, bool, error) {
	newLimit := paging.Limit + 1
	var items []models.Item
	query := sqlDB.DB.Model(&models.Item{}).Limit(newLimit).Offset(paging.Offset)
	if paging.Type != "" {
		query.Where("type = ?", paging.Type)
	}
	result := query.Find(&items)
	hasMore := false
	if len(items) == newLimit {
		hasMore = true
	}
	if len(items) <= paging.Limit {
		return items, hasMore, result.Error
	}

	return items[0:paging.Limit], hasMore, result.Error
}
