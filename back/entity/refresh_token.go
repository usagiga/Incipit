package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

// RefreshToken represents security token used to authorize admin user
type RefreshToken struct {
	gorm.Model
	Token     string
	ExpiredAt time.Time

	AdminUserID uint
	AdminUser   AdminUser
}

func (rt *RefreshToken) IsExpired(time time.Time) (isExpired bool) {
	return rt.ExpiredAt.Before(time)
}
