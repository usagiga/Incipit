package model

import (
	"github.com/usagiga/Incipit/back/entity"
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"golang.org/x/xerrors"
	"regexp"
)

type AdminUserValidatorImpl struct {
	unusableCharRule *regexp.Regexp
}

func NewAdminUserValidator() (auv AdminUserValidator, err error) {
	usableCharRule, err := regexp.Compile("[^a-zA-Z0-9_]")
	if err != nil {
		return nil, err
	}

	return &AdminUserValidatorImpl{unusableCharRule: usableCharRule}, nil
}

func (m *AdminUserValidatorImpl) ValidateAll(user *entity.AdminUser) (err error) {
	errChan := make(chan error, 4)

	go func() {
		errChan <- m.ValidateID(user.ID)
		errChan <- m.ValidateName(user.Name)
		errChan <- m.ValidateScreenName(user.ScreenName)
		errChan <- m.ValidatePassword(user.Password)
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return xerrors.Errorf("This user is invalid: %w", err)
		}
	}

	return nil
}

func (m *AdminUserValidatorImpl) ValidateID(id uint) (err error) {
	return nil
}

func (m *AdminUserValidatorImpl) ValidateName(name string) (err error) {
	errChan := make(chan error, 3)
	nameBytes := []byte(name)

	go func() {
		if len(name) < 3 {
			errChan <- interr.NewDistinctError("Name is too short for admin user", interr.AdminUserValidation, interr.AdminUserValidation_NameIsTooShort, nil)
		}
		if 32 < len(name) {
			errChan <- interr.NewDistinctError("Name is too long for admin user", interr.AdminUserValidation, interr.AdminUserValidation_NameIsTooLong, nil)
		}
		if m.unusableCharRule.Match(nameBytes) {
			errChan <- interr.NewDistinctError("Name has unavailable char for admin user", interr.AdminUserValidation, interr.AdminUserValidation_NameHasUnavailableChar, nil)
		}

		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return xerrors.Errorf("This name for user is invalid: %w", err)
		}
	}

	return nil
}

func (m *AdminUserValidatorImpl) ValidateScreenName(scName string) (err error) {
	errChan := make(chan error, 3)
	scNameBytes := []byte(scName)

	go func() {
		if len(scName) < 3 {
			errChan <- interr.NewDistinctError("Screen name is too short for admin user", interr.AdminUserValidation, interr.AdminUserValidation_ScreenNameIsTooShort, nil)
		}
		if 32 < len(scName) {
			errChan <- interr.NewDistinctError("Screen name is too long for admin user", interr.AdminUserValidation, interr.AdminUserValidation_ScreenNameIsTooLong, nil)
		}
		if m.unusableCharRule.Match(scNameBytes) {
			errChan <- interr.NewDistinctError("Screen name has unavailable char for admin user", interr.AdminUserValidation, interr.AdminUserValidation_ScreenNameHasUnavailableChar, nil)
		}

		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return xerrors.Errorf("This screen name for user is invalid: %w", err)
		}
	}

	return nil
}

func (m *AdminUserValidatorImpl) ValidatePassword(password string) (err error) {
	errChan := make(chan error, 3)
	passwordBytes := []byte(password)

	go func() {
		if len(password) < 8 {
			errChan <- interr.NewDistinctError("Password is too short for admin user", interr.AdminUserValidation, interr.AdminUserValidation_PasswordIsTooShort, nil)
		}
		if 72 < len(password) {
			errChan <- interr.NewDistinctError("Password is too long for admin user", interr.AdminUserValidation, interr.AdminUserValidation_PasswordIsTooLong, nil)
		}
		if m.unusableCharRule.Match(passwordBytes) {
			errChan <- interr.NewDistinctError("Password has unavailable char for admin user", interr.AdminUserValidation, interr.AdminUserValidation_PasswordHasUnavailableChar, nil)
		}

		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return xerrors.Errorf("Password for user is invalid: %w", err)
		}
	}

	return nil
}
