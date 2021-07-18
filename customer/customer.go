package customer

import (
	"gorm.io/gorm"
	"time"
)

const (
	Street = iota
	Avenue
	Road
)

type Customer struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Address Address `json:"address"`
	Orders []int64 `json:"order"`

	CreateAt time.Time `json:"created"`
	UpdateAt time.Time `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type Address struct {
	ID uint `json:"id" gorm:"primaryKey"`
	IDUser uint `json:"userID"`
	Type int64 `json:"type"`
	PublicPlace string `json:"publicPlace"`
	Number int64 `json:"number"`
	Complement string `json:"complement"`

	CreateAt time.Time `json:"created"`
	UpdateAt time.Time `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type LoginUser struct {
	ID uint `json:"id" gorm:"primaryKey"`
	IDUser uint `json:"userID"`
	User string `json:"Login"`
	Password string `json:"password"`
}