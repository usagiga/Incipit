package testutils

import (
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
)

func DropTestDB(db *gorm.DB) (err error) {
	errChan := make(chan error, 5)

	go func() {
		errChan <- db.DropTableIfExists(&entity.Link{}).Error
		errChan <- db.DropTableIfExists(&entity.AccessToken{}).Error
		errChan <- db.DropTableIfExists(&entity.RefreshToken{}).Error
		errChan <- db.DropTableIfExists(&entity.AdminUser{}).Error
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
