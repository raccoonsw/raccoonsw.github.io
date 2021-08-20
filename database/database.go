package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"restApiProject/models"
)

//func (SubItem) TableName() string {
//	return "items"
//}

func Connect(config models.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect models")
	}

	//userFlag := db.Migrator().h.HasTable(&User{})
	itemFlag := db.Migrator().HasTable(&models.Item{})
	if !itemFlag {
		err = db.Migrator().CreateTable(&models.Item{})
		if err != nil {
			panic("failed to create Items table in models")
		}
	}

	orderFlag := db.Migrator().HasTable(&models.Order{})
	if !orderFlag {
		err = db.Migrator().CreateTable(&models.Order{})
		if err != nil {
			panic("failed to create Items table in models")
		}
	}
	return db
}

func (sqlDB *DBModel) ClearTable() {
	itemFlag := sqlDB.DB.Migrator().HasTable(&models.Item{})
	if itemFlag {
		err := sqlDB.DB.Migrator().DropTable(&models.Item{})
		err = sqlDB.DB.Migrator().CreateTable(&models.Item{})
		if err != nil {
			panic("failed to create Items table in models")
		}
	}
}

func (sqlDB *DBModel) Close() {
	db, err := sqlDB.DB.DB()
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}
