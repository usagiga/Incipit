package entity

import "github.com/jinzhu/gorm"

// AdminUser represents Incipit administrator
type AdminUser struct {
	gorm.Model
	Name string
	ScreenName string
	Password string
}
