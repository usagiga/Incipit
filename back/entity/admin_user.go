package entity

import "github.com/jinzhu/gorm"

// AdminUser represents Incipit administrator
type AdminUser struct {
	gorm.Model
	Name string `gorm:"unique"`
	ScreenName string
	Password string
}
