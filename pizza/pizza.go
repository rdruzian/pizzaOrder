package pizza

import (
	"gorm.io/gorm"
	"time"
)

type Pizza struct {
	ID uint `json:"id" gorm:"primaryKey"`
	FlavorName  string   `json:"flavorName"`
	Ingredients []string `json:"ingredients"`
	Price       float64  `json:"price"`

	CreateAt time.Time `json:"created"`
	UpdateAt time.Time `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
