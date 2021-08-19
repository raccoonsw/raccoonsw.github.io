package models

//type User struct {
//	Id	int	`gorm:"primaryKey"`
//	FullName	string	`gorm:"unique;not null;type:varchar(100)"`
//}

type Item struct {
	Id   int     `gorm:"primaryKey" json:"id" binding:"omitempty,gt=0"`
	Sku  string  `gorm:"unique;not null;type:varchar(100)" json:"sku" binding:"required,ascii,lowercase,lte=100"`
	Name string  `gorm:"not null;type:varchar(100)" json:"name" binding:"required,lte=100,ascii"`
	Type string  `gorm:"not null;type:varchar(100)" json:"type" binding:"required,oneof=virtual_good virtual_currency bundle"`
	Cost float32 `gorm:"not null" json:"cost" binding:"required"`
	//UserId	int	//`gorm:"not null"`
}
