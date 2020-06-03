package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
)

type AdminModelImpl struct {
	db *gorm.DB
}

func NewAdminModel(db *gorm.DB) AdminModel {
	return &AdminModelImpl{
		db: db,
	}
}

func (m *AdminModelImpl) Add(user *entity.AdminUser) (added *entity.AdminUser, err error) {
	// Add the row
	result := m.db.Create(user)
	err = result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, errors.New("AdminModel.Add(): Can't add it")
	}

	return user, err
}

func (m *AdminModelImpl) FindOne(id uint) (user *entity.AdminUser, err error) {
	user = &entity.AdminUser{}

	// Find the row
	result := m.db.First(user, id)
	err = result.Error
	if result.RecordNotFound() {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, err
}

func (m *AdminModelImpl) FindOneByName(name string) (user *entity.AdminUser, err error) {
	user = &entity.AdminUser{}
	condition := &entity.AdminUser{Name: name}

	// Find the row with the condition
	result := m.db.Where(condition).First(user)
	err = result.Error
	if result.RecordNotFound() {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, err
}

func (m *AdminModelImpl) Find() (users []entity.AdminUser, err error) {
	// Find all rows
	result := m.db.Find(&users)
	err = result.Error
	if result.RecordNotFound() {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return users, err
}

func (m *AdminModelImpl) Update(updating *entity.AdminUser) (updated *entity.AdminUser, err error) {
	// Finding the row
	found, err := m.FindOne(updating.ID)
	if err != nil {
		return nil, err
	}
	if found == nil {
		return nil, errors.New("AdminModel.Update(): Not found updating row")
	}

	// Update the row
	result := m.db.Model(&entity.AdminUser{}).Update(updating)
	err = result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return updating, nil
	}

	return updating, err
}

func (m *AdminModelImpl) Delete(id uint) (err error) {
	// Delete the row
	result := m.db.Unscoped().Delete(&entity.AdminUser{}, id)
	err = result.Error
	if err != nil {
		return nil
	}
	if result.RowsAffected <= 0 {
		return errors.New("AdminModel.Delete(): Can't delete it")
	}

	return nil
}
