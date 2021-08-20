package models

import "time"

type Order struct {
	Id        int    `gorm:"primaryKey" json:"id" binding:"omitempty,gt=0"`
	ItemId    int    `gorm:"not null;type:int" json:"item_id" binding:"required,gt=0"`
	Email     string `gorm:"not null;type:varchar(100)" json:"email" binding:"required,email"`
	CreatedAt time.Time
	//UserId	int	//`gorm:"not null"`
}
