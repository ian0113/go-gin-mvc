package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}
