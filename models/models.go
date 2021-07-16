package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBModel struct {
	DB *gorm.DB
}

type Config struct {
	User     string
	Password string
	DBName   string
}

func Connect(config Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", config.User, config.Password, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect models")
	}

	//userFlag := db.Migrator().h.HasTable(&User{})
	itemFlag := db.Migrator().HasTable(&Item{})
	if !itemFlag {
		////err = sqlDB.AutoMigrate(&User{}, &Item{})
		//err = sqlDB.Migrator().CreateTable(&User{})
		//if err != nil {
		//	panic("failed to create Users table in models")
		//}
		err = db.Migrator().CreateTable(&Item{})
		if err != nil {
			panic("failed to create Items table in models")
		}
	}
	return db
}

func (sqlDB DBModel) ClearTable() {
	itemFlag := sqlDB.DB.Migrator().HasTable(&Item{})
	if itemFlag {
		err := sqlDB.DB.Migrator().DropTable(&Item{})
		err = sqlDB.DB.Migrator().CreateTable(&Item{})
		if err != nil {
			panic("failed to create Items table in models")
		}
	}
}

func (sqlDB DBModel) Close() {
	db, err := sqlDB.DB.DB()
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}

//type User struct {
//	Id	uint	`gorm:"primaryKey"`
//	FullName	string	`gorm:"unique;not null;type:varchar(100)"`
//}

type Item struct {
	Id   uint    `gorm:"primaryKey" json:"id"` //;autoIncrement
	Sku  string  `gorm:"unique;not null;type:varchar(100)" json:"sku" binding:"required,ascii,lowercase,lte=100"`
	Name string  `gorm:"not null;type:varchar(100)" json:"name" binding:"required,lte=100,ascii"`
	Type string  `gorm:"not null;type:varchar(100)" json:"type" binding:"required,oneof=virtual_good virtual_currency bundle"`
	Cost float32 `gorm:"not null" json:"cost" binding:"required"`
	//UserId	uint	//`gorm:"not null"`
}

//func (SubItem) TableName() string {
//	return "items"
//}

func (sqlDB *DBModel) CreateItem(item Item) (Item, error) {
	result := sqlDB.DB.Create(&item)
	return item, result.Error
}

func (sqlDB *DBModel) UpdateItemById(item Item) error {
	result := sqlDB.DB.Save(&item)
	return result.Error
}

func (sqlDB *DBModel) UpdateItemBySku(item Item) error {
	result := sqlDB.DB.Model(&Item{}).Where("sku = ?", item.Sku).Updates(item)
	return result.Error
}

func (sqlDB *DBModel) DeleteItemById(id uint) error {
	result := sqlDB.DB.Delete(&Item{}, id)
	return result.Error
}

func (sqlDB *DBModel) DeleteItemBySku(sku string) error {
	result := sqlDB.DB.Where("sku = ?", sku).Delete(&Item{})
	return result.Error
}

func (sqlDB *DBModel) GetItemById(id uint) (Item, error) {
	var newItem Item
	result := sqlDB.DB.First(&newItem, id)
	return newItem, result.Error
}

func (sqlDB *DBModel) GetItemBySku(sku string) (Item, error) {
	var newItem Item
	result := sqlDB.DB.First(&newItem, "sku = ?", sku)
	return newItem, result.Error
}

func (sqlDB *DBModel) GetAllItems(limit int, offset int) ([]Item, bool, error) {
	newLimit := limit + 1
	var items []Item
	result := sqlDB.DB.Model(&Item{}).Limit(newLimit).Offset(offset).Find(&items)
	hasMore := false
	if len(items) == newLimit {
		hasMore = true
	}
	return items[0:limit], hasMore, result.Error
}