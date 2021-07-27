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
	Type uint `json:"type"`
	PublicPlace string `json:"publicPlace"`
	Number int64 `json:"number"`
	Complement string `json:"complement"`
	Login string `json:"login"`
	Password string `json:"password"`

	CreateAt time.Time `json:"created"`
	UpdateAt time.Time `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}