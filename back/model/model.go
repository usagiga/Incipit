package model

import "github.com/usagiga/Incipit/back/entity"

// InstallerModel treats starting to use Incipit.
// If there's no admin, this will be called.
type InstallerModel interface {
	CreateNewAdmin(name, screenName, password string) (
		accToken *entity.AccessToken,
		refToken *entity.RefreshToken,
		err error,
	)
	IsNeededInstall() (isNeeded bool, err error)
}

// AdminModel treats CRUD of admin user.
type AdminModel interface {
	Add(user *entity.AdminUser) (added *entity.AdminUser, err error)
	FindOne(id uint) (user *entity.AdminUser, err error)
	FindOneByName(name string) (user *entity.AdminUser, err error)
	Find() (users []entity.AdminUser, err error)
	Update(updating *entity.AdminUser) (updated *entity.AdminUser, err error)
	Delete(id uint) (err error)
}

// AdminAuthModel treats authentication for admin user.
type AdminAuthModel interface {
	Authorize(accTokenStr string) (user *entity.AdminUser, err error)
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
	FindOneByShortID(shortId string) (link *entity.Link, err error)
	Find() (links []entity.Link, err error)
	Update(updating *entity.Link) (updated *entity.Link, err error)
	Delete(id uint) (err error)
}

// HashModel treats hashing and validating for password.
type HashModel interface {
	Generate(pass string) (hashed string, err error)
	Equals(hashedPassword, rawPassword string) (err error)
}

// AdminUserValidator treats validation admin user in request arguments
type AdminUserValidator interface {
	ValidateAll(user *entity.AdminUser) (err error)
	ValidateID(id uint) (err error)
	ValidateName(name string) (err error)
	ValidateScreenName(scName string) (err error)
	ValidatePassword(password string) (err error)
}

// LinkValidator treats validation link in request arguments
type LinkValidator interface {
	ValidateAll(link *entity.Link) (err error)
	ValidateID(id uint) (err error)
	ValidateURL(url string) (err error)
}
