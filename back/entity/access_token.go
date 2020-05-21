package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

// AccessToken represents security token used to authorize admin user
type AccessToken struct {
	gorm.Model
	Token string
	ExpiredAt *time.Time

	AdminUserID uint
	AdminUser AdminUser
}
