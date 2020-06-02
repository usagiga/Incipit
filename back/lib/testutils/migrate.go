package testutils

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
)

func MigrateTestDB(db *gorm.DB) (err error) {
	linkObj := &entity.Link{}
	adminUsrObj := &entity.AdminUser{}
	accTokenObj := &entity.AccessToken{}
	refTokenObj := &entity.RefreshToken{}

	// Auto migrate
	err = db.
		Set("gorm:table_options", "CHARSET=utf8mb4").
		AutoMigrate(
			linkObj,
			adminUsrObj,
			accTokenObj,
			refTokenObj,
		).
		Error

	if err != nil {
		return errors.New("main.Migrate(): Can't migrate automatically")
	}

	// Build their relationship
	err = db.
		Model(accTokenObj).
		AddForeignKey(
			"admin_user_id",
			"admin_users(id)",
			"CASCADE",
			"CASCADE",
		).
		Error

	if err != nil {
		return errors.New("main.Migrate(): Can't define relationship between AdminUser and AccessToken")
	}

	err = db.
		Model(refTokenObj).
		AddForeignKey(
			"admin_user_id",
			"admin_users(id)",
			"CASCADE",
			"CASCADE",
		).
		Error
	if err != nil {
		return errors.New("main.Migrate(): Can't define relationship between AdminUser and RefreshToken")
	}

	return nil
}
