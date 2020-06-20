package model

import (
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"golang.org/x/xerrors"
)

type AdminModelImpl struct {
	db                 *gorm.DB
	hashModel          HashModel
	adminUserValidator AdminUserValidator
}

func NewAdminModel(
	db *gorm.DB,
	hashModel HashModel,
	adminUserValidator AdminUserValidator,
) AdminModel {
	return &AdminModelImpl{
		db:                 db,
		hashModel:          hashModel,
		adminUserValidator: adminUserValidator,
	}
}

func (m *AdminModelImpl) Add(user *entity.AdminUser) (added *entity.AdminUser, err error) {
	// Validate args
	err = m.adminUserValidator.ValidateAll(user)
	if err != nil {
		return nil, xerrors.Errorf("It was passed bad arguments: %w", err)
	}

	// Generate hash of password
	user.Password, err = m.hashModel.Generate(user.Password)
	if err != nil {
		return nil, xerrors.Errorf("Can't hash password: %w", err)
	}

	// Add the row
	result := m.db.Create(user)
	err = result.Error
	if err != nil || result.RowsAffected <= 0 {
		return nil, interr.NewDistinctError("Can't add user", interr.AdminModel, interr.AdminModel_FailedAdd, nil).Wrap(err)
	}

	return user, nil
}

func (m *AdminModelImpl) FindOne(id uint) (user *entity.AdminUser, err error) {
	user = &entity.AdminUser{}

	// Validate args
	err = m.adminUserValidator.ValidateID(id)
	if err != nil {
		return nil, xerrors.Errorf("It was passed bad arguments: %w", err)
	}

	// Find the row
	result := m.db.First(user, id)
	err = result.Error
	if result.RecordNotFound() {
		return nil, interr.NewDistinctError("There's no user", interr.AdminModel, interr.AdminModel_FindingUserNotFound, nil).Wrap(err)
	}
	if err != nil {
		return nil, interr.NewDistinctError("Can't find user", interr.AdminModel, interr.AdminModel_FailedFind, nil).Wrap(err)
	}

	return user, nil
}

func (m *AdminModelImpl) FindOneByName(name string) (user *entity.AdminUser, err error) {
	user = &entity.AdminUser{}
	condition := &entity.AdminUser{Name: name}

	// Validate args
	err = m.adminUserValidator.ValidateName(name)
	if err != nil {
		return nil, xerrors.Errorf("It was passed bad arguments: %w", err)
	}

	// Find the row with the condition
	result := m.db.Where(condition).First(user)
	err = result.Error
	if result.RecordNotFound() {
		return nil, interr.NewDistinctError("There's no user", interr.AdminModel, interr.AdminModel_FindingUserNotFound, nil).Wrap(err)
	}
	if err != nil {
		return nil, interr.NewDistinctError("Can't find user", interr.AdminModel, interr.AdminModel_FailedFind, nil).Wrap(err)
	}

	return user, nil
}

func (m *AdminModelImpl) Find() (users []entity.AdminUser, err error) {
	// Find all rows
	result := m.db.Find(&users)
	err = result.Error
	if result.RecordNotFound() {
		return nil, interr.NewDistinctError("There's no user", interr.AdminModel, interr.AdminModel_FindingUserNotFound, nil).Wrap(err)
	}
	if err != nil {
		return nil, interr.NewDistinctError("Can't find user", interr.AdminModel, interr.AdminModel_FailedFind, nil).Wrap(err)
	}

	return users, nil
}

func (m *AdminModelImpl) Update(updating *entity.AdminUser) (updated *entity.AdminUser, err error) {
	// Validate args
	err = m.adminUserValidator.ValidateAll(updating)
	if err != nil {
		return nil, xerrors.Errorf("It was passed bad arguments: %w", err)
	}

	// Finding the row
	_, err = m.FindOne(updating.ID)
	if err != nil {
		return nil, interr.NewDistinctError("Can't find user", interr.AdminModel, interr.AdminModel_UpdatingUserNotFound, nil).Wrap(err)
	}

	// If password isn't default value, generate hash of password
	if updating.Password != "" {
		updating.Password, err = m.hashModel.Generate(updating.Password)
		if err != nil {
			return nil, xerrors.Errorf("Can't hash password: %w", err)
		}
	}

	// Update the row
	result := m.db.Model(&entity.AdminUser{}).Update(updating)
	err = result.Error
	if err != nil {
		return nil, interr.NewDistinctError("Can't update user", interr.AdminModel, interr.AdminModel_FailedUpdate, nil).Wrap(err)
	}
	if result.RowsAffected <= 0 {
		// There's no value to change
		return updating, nil
	}

	return updating, nil
}

func (m *AdminModelImpl) Delete(id uint) (err error) {
	// Validate args
	err = m.adminUserValidator.ValidateID(id)
	if err != nil {
		return xerrors.Errorf("It was passed bad arguments: %w", err)
	}

	// Delete the row
	result := m.db.Unscoped().Delete(&entity.AdminUser{}, id)
	err = result.Error
	if err != nil {
		return interr.NewDistinctError("Can't delete user", interr.AdminModel, interr.AdminModel_FailedDelete, nil)
	}
	if result.RowsAffected <= 0 {
		return interr.NewDistinctError("Can't delete user", interr.AdminModel, interr.AdminModel_FailedDelete, nil).Wrap(err)
	}

	return nil
}
