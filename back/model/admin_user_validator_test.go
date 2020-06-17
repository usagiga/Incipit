package model

import (
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/testutils"
	"testing"
)

func initAdminUserValidator(t *testing.T) (av AdminUserValidator, finalizer testutils.Finalizer) {
	av, err := NewAdminUserValidator()
	if err != nil {
		t.Fatalf("Can't initialize AdminUserValidator: %v", err)
	}

	finalizer = func() {}

	return av, finalizer
}

func TestAdminUserValidatorImpl_ValidateAll(t *testing.T) {
	// Initialize model
	av, avFin := initAdminUserValidator(t)
	defer avFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError  bool
		TestingAdminUser *entity.AdminUser
	}{
		// Test for Name
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "", ScreenName: "screen_1", Password: "password_1"}},
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "12", ScreenName: "screen_1", Password: "password_1"}},                                // 2 chars
		{IsExpectedError: false, TestingAdminUser: &entity.AdminUser{Name: "123", ScreenName: "screen_1", Password: "password_1"}},                              // 3 chars
		{IsExpectedError: false, TestingAdminUser: &entity.AdminUser{Name: "12345678901234567890123456789012", ScreenName: "screen_1", Password: "password_1"}}, // 32 chars
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "123456789012345678901234567890123", ScreenName: "screen_1", Password: "password_1"}}, // 33 chars
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "Unusable Name!", ScreenName: "screen_1", Password: "password_1"}},
		{IsExpectedError: false, TestingAdminUser: &entity.AdminUser{Name: "_____", ScreenName: "screen_1", Password: "password_1"}},

		// Test for ScreenName
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "", Password: "password_1"}},
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "12", Password: "password_1"}},                                // 2 chars
		{IsExpectedError: false, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "123", Password: "password_1"}},                              // 3 chars
		{IsExpectedError: false, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "12345678901234567890123456789012", Password: "password_1"}}, // 32 chars
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "123456789012345678901234567890123", Password: "password_1"}}, // 33 chars
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "Unusable Screen Name!", Password: "password_1"}},
		{IsExpectedError: false, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "_____", Password: "password_1"}},

		// Test for Password
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: ""}},
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "1234567"}},                                                                   // 7 chars
		{IsExpectedError: false, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "12345678"}},                                                                 // 8 chars
		{IsExpectedError: false, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "123456789012345678901234567890123456789012345678901234567890123456789012"}}, // 72 chars
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "1234567890123456789012345678901234567890123456789012345678901234567890123"}}, // 73 chars
		{IsExpectedError: true, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "Unusable Password!"}},
		{IsExpectedError: false, TestingAdminUser: &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "________"}},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		testingUser := v.TestingAdminUser

		err := av.ValidateAll(testingUser)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}
	}
}

func TestAdminUserValidatorImpl_ValidateID(t *testing.T) {
	// NOP
}

func TestAdminUserValidatorImpl_ValidateName(t *testing.T) {
	// Initialize model
	av, avFin := initAdminUserValidator(t)
	defer avFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingName     string
	}{
		{IsExpectedError: true, TestingName: ""},
		{IsExpectedError: true, TestingName: "12"},                                // 2 chars
		{IsExpectedError: false, TestingName: "123"},                              // 3 chars
		{IsExpectedError: false, TestingName: "12345678901234567890123456789012"}, // 32 chars
		{IsExpectedError: true, TestingName: "123456789012345678901234567890123"}, // 33 chars
		{IsExpectedError: true, TestingName: "Unusable Screen Name!"},
		{IsExpectedError: false, TestingName: "_____"},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		testingName := v.TestingName

		err := av.ValidateName(testingName)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}
	}
}

func TestAdminUserValidatorImpl_ValidateScreenName(t *testing.T) {
	// Initialize model
	av, avFin := initAdminUserValidator(t)
	defer avFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError   bool
		TestingScreenName string
	}{
		{IsExpectedError: true, TestingScreenName: ""},
		{IsExpectedError: true, TestingScreenName: "12"},                                // 2 chars
		{IsExpectedError: false, TestingScreenName: "123"},                              // 3 chars
		{IsExpectedError: false, TestingScreenName: "12345678901234567890123456789012"}, // 32 chars
		{IsExpectedError: true, TestingScreenName: "123456789012345678901234567890123"}, // 33 chars
		{IsExpectedError: true, TestingScreenName: "Unusable Screen Name!"},
		{IsExpectedError: false, TestingScreenName: "_____"},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		testingName := v.TestingScreenName

		err := av.ValidateScreenName(testingName)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}
	}
}

func TestAdminUserValidatorImpl_ValidatePassword(t *testing.T) {
	// Initialize model
	av, avFin := initAdminUserValidator(t)
	defer avFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingPassword string
	}{
		{IsExpectedError: true, TestingPassword: ""},
		{IsExpectedError: true, TestingPassword: "1234567"},                                                                   // 7 chars
		{IsExpectedError: false, TestingPassword: "12345678"},                                                                 // 8 chars
		{IsExpectedError: false, TestingPassword: "123456789012345678901234567890123456789012345678901234567890123456789012"}, // 72 chars
		{IsExpectedError: true, TestingPassword: "1234567890123456789012345678901234567890123456789012345678901234567890123"}, // 73 chars
		{IsExpectedError: true, TestingPassword: "Unusable Screen Name!"},
		{IsExpectedError: false, TestingPassword: "__________"},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		testingPassword := v.TestingPassword

		err := av.ValidatePassword(testingPassword)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}
	}
}
