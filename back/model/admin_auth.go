package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/usagiga/Incipit/back/entity"
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"golang.org/x/xerrors"
	"time"
)

type AdminAuthModelImpl struct {
	db         *gorm.DB
	adminModel AdminModel
	hashModel  HashModel
}

func NewAdminAuthModel(
	db *gorm.DB,
	adminModel AdminModel,
	hashModel HashModel,
) AdminAuthModel {
	return &AdminAuthModelImpl{
		db:         db,
		adminModel: adminModel,
		hashModel:  hashModel,
	}
}

func (m *AdminAuthModelImpl) Authorize(accTokenStr string) (user *entity.AdminUser, err error) {
	foundToken := &entity.AccessToken{}
	now := time.Now()

	// Find specified user
	result := m.db.Preload("AdminUser").Where(&entity.AccessToken{Token: accTokenStr}).First(foundToken)
	err = result.Error
	if result.RecordNotFound() {
		return nil, interr.NewDistinctError("There's no token", interr.AdminAuthModel, interr.AdminAuthModel_FailedToFindUser, nil).Wrap(err)
	}
	if err != nil {
		return nil, interr.NewDistinctError("Can't find token", interr.AdminAuthModel, interr.AdminAuthModel_FailedToFindUser, nil).Wrap(err)
	}

	// Check it is expired
	if foundToken.IsExpired(now) {
		return nil, interr.NewDistinctError("This token is expired", interr.AdminAuthModel, interr.AdminAuthModel_ExpiredToken, nil)
	}

	return &foundToken.AdminUser, nil
}

func (m *AdminAuthModelImpl) Login(name, password string) (accToken *entity.AccessToken, refToken *entity.RefreshToken, err error) {
	// Find specified user
	user, err := m.adminModel.FindOneByName(name)
	if err != nil {
		return nil, nil, interr.NewDistinctError("Can't find user by token", interr.AdminAuthModel, interr.AdminAuthModel_FailedToFindUser, nil).Wrap(err)
	}
	if user == nil {
		return nil, nil, interr.NewDistinctError("There's no user", interr.AdminAuthModel, interr.AdminAuthModel_FailedToFindUser, nil)
	}

	// Check its password equals specified password
	err = m.hashModel.Equals(user.Password, password)
	if err != nil {
		return nil, nil, interr.NewDistinctError("Don't match password", interr.AdminAuthModel, interr.AdminAuthModel_UnmatchPassword, nil).Wrap(err)
	}

	// Generate Access / Refresh Token
	accToken, refToken = m.generateTokenPair(user.ID)
	err = m.saveTokenPair(accToken, refToken)
	if err != nil {
		return nil, nil, interr.NewDistinctError("Can't save token pair", interr.AdminAuthModel, interr.AdminAuthModel_FailedToStoreToken, nil).Wrap(err)
	}

	return accToken, refToken, nil
}

func (m *AdminAuthModelImpl) RenewAccessToken(refTokenStr string) (accToken *entity.AccessToken, refToken *entity.RefreshToken, err error) {
	foundToken := &entity.RefreshToken{}
	now := time.Now()

	// Find specified user
	result := m.db.Preload("AdminUser").Where(&entity.RefreshToken{Token: refTokenStr}).First(foundToken)
	err = result.Error
	if result.RecordNotFound() {
		return nil, nil, interr.NewDistinctError("There's no token", interr.AdminAuthModel, interr.AdminAuthModel_FailedToFindUser, nil).Wrap(err)
	}
	if err != nil {
		return nil, nil, interr.NewDistinctError("Can't find token", interr.AdminAuthModel, interr.AdminAuthModel_FailedToFindUser, nil).Wrap(err)
	}

	// Check it is expired
	if foundToken.IsExpired(now) {
		return nil, nil, interr.NewDistinctError("This token is expired", interr.AdminAuthModel, interr.AdminAuthModel_ExpiredToken, nil)
	}

	// Generate access / refresh token
	accToken, refToken = m.generateTokenPair(foundToken.AdminUserID)
	err = m.saveTokenPair(accToken, refToken)
	if err != nil {
		return nil, nil, interr.NewDistinctError("Can't save token pair", interr.AdminAuthModel, interr.AdminAuthModel_FailedToStoreToken, nil).Wrap(err)
	}

	return accToken, refToken, nil
}

func (m *AdminAuthModelImpl) generateTokenPair(userId uint) (accToken *entity.AccessToken, refToken *entity.RefreshToken) {
	now := time.Now()
	accTokenStr := uuid.NewV4().String()
	accTokenExp := now.AddDate(0, 0, 1)
	refTokenStr := uuid.NewV4().String()
	refTokenExp := now.AddDate(0, 0, 14)

	accToken = &entity.AccessToken{
		Token:       accTokenStr,
		ExpiredAt:   accTokenExp,
		AdminUserID: userId,
	}

	refToken = &entity.RefreshToken{
		Token:       refTokenStr,
		ExpiredAt:   refTokenExp,
		AdminUserID: userId,
	}

	return accToken, refToken
}

func (m *AdminAuthModelImpl) saveTokenPair(accToken *entity.AccessToken, refToken *entity.RefreshToken) (err error) {
	tx := m.db.Begin()
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	err = tx.Error
	if err != nil {
		return xerrors.Errorf("Can't begin transaction: %w", err)
	}

	err = tx.Create(accToken).Error
	if err != nil {
		tx.Rollback()
		return xerrors.Errorf("Can't add access token: %w", err)
	}

	err = tx.Create(refToken).Error
	if err != nil {
		tx.Rollback()
		return xerrors.Errorf("Can't add refresh token: %w", err)
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return xerrors.Errorf("Can't commit transaction: %w", err)
	}

	return nil
}
