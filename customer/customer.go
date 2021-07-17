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
	Favorite string `json:"favorite"`
	Login string `json:"login"`
	Password string `json:"pass"`

	CreateAt time.Time `json:"created"`
	UpdateAt time.Time `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type Address struct {
	Type int64 `json:"type"`
	PublicPlace string `json:"publicPlace"`
	Number int64 `json:"number"`
	Complement string `json:"complement"`

	CreateAt time.Time `json:"created"`
	UpdateAt time.Time `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}