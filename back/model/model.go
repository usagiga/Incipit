package model

import "github.com/usagiga/Incipit/back/entity"

// InstallerModel treats starting to use Incipit.
// If there's no admin, this will be called.
type InstallerModel interface {
	CreateNewAdmin(email, password string) (
		accToken *entity.AccessToken,
		refToken *entity.RefreshToken,
		err error,
	)
}

// AdminModel treats CRUD of admin user.
type AdminModel interface {
	Add(user *entity.AdminUser) (added *entity.AdminUser, err error)
	FindOne(id uint) (user *entity.AdminUser, err error)
	Find() (users []entity.AdminUser, err error)
	Update(updating *entity.AdminUser) (updated *entity.AdminUser, err error)
	Delete(id uint) (err error)
}

// AdminAuthModel treats authentication for admin user.
type AdminAuthModel interface {
	Authorize(accToken *entity.AccessToken) (user *entity.AdminUser, err error)
	Login(name, password string) (
		accToken *entity.AccessToken,
		refToken *entity.RefreshToken,
		err error,
	)
	RenewAccessToken(refTokenStr string) (
		accToken *entity.AccessToken,
		refToken *entity.RefreshToken,
		err error,
	)
}

// LinkModel treats CRUD of shortened link.
type LinkModel interface {
	Add(link *entity.Link) (added *entity.Link, err error)
	FindOne(id uint) (link *entity.Link, err error)
	Find() (links []entity.Link, err error)
	Update(updating *entity.Link) (updated *entity.Link, err error)
	Delete(id uint) (err error)
}
