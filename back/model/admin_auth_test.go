package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/testutils"
	"testing"
	"time"
)

func initAdminAuthModel(t *testing.T) (aam AdminAuthModel, finalizer testutils.Finalizer) {
	errChan := make(chan error, 3)
	db, dbFin := testutils.SetupTestDB(t)

	// Insert test values
	go func() {
		now := time.Now()

		errChan <- db.Create(&entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "hashed_password_1"}).Error
		errChan <- db.Create(&entity.AdminUser{Name: "user_2", ScreenName: "screen_2", Password: "hashed_password_2"}).Error

		errChan <- db.Create(&entity.AccessToken{Token: "111", ExpiredAt: now.AddDate(1, 0, 0), AdminUserID: 1}).Error
		errChan <- db.Create(&entity.AccessToken{Token: "222", ExpiredAt: now.AddDate(-1, 0, 0), AdminUserID: 2}).Error

		errChan <- db.Create(&entity.RefreshToken{Token: "111", ExpiredAt: now.AddDate(1, 0, 0), AdminUserID: 1}).Error
		errChan <- db.Create(&entity.RefreshToken{Token: "222", ExpiredAt: now.AddDate(-1, 0, 0), AdminUserID: 2}).Error

		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			t.Fatal("initAdminAuthModel(): Can't insert testing values; ", err)
		}
	}

	return &AdminAuthModelImpl{db: db, adminModel: &adminModelStub{}, hashModel: &hashModelStub{}}, dbFin
}

func TestAdminAuthModelImpl_Authorize(t *testing.T) {
	// Initialize model
	aam, aamFin := initAdminAuthModel(t)
	defer aamFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingTokenStr string
		ExpectingValue  *entity.AdminUser
	}{
		{IsExpectedError: false, TestingTokenStr: "111", ExpectingValue: &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "hashed_password_1", Model: gorm.Model{ID: 1}}},
		{IsExpectedError: true, TestingTokenStr: "222", ExpectingValue: nil},
		{IsExpectedError: true, TestingTokenStr: "777", ExpectingValue: nil},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		authTokenStr := v.TestingTokenStr
		expectedVal := v.ExpectingValue

		actualVal, err := aam.Authorize(authTokenStr)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When actual value isn't equal expected value
		if !adminUserEquals(expectedVal, actualVal) {
			t.Fatalf("Case %d: Actual value isn't equal expected value.\nExpected:\t%v,\nActual:\t%v", caseNum, expectedVal, actualVal)
		}
	}
}

func TestAdminAuthModelImpl_Login(t *testing.T) {
	// Initialize model
	aam, aamFin := initAdminAuthModel(t)
	defer aamFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingUserName string
		TestingPassword string
		ExpectedUserID  uint
	}{
		{IsExpectedError: false, TestingUserName: "user_1", TestingPassword: "password_1", ExpectedUserID: 1},
		{IsExpectedError: true, TestingUserName: "user_2", TestingPassword: "error_password"},
		{IsExpectedError: true, TestingUserName: "error_user", TestingPassword: "error_password"},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		testUser := v.TestingUserName
		testPass := v.TestingPassword
		expectedUserID := v.ExpectedUserID

		accToken, refToken, err := aam.Login(testUser, testPass)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When actual value isn't equal expected value
		if isExpectedError && accToken != nil {
			t.Fatalf("Case %d: Generated access token isn't nil", caseNum)
		}
		if isExpectedError && accToken != nil {
			t.Fatalf("Case %d: Generated access token isn't nil", caseNum)
		}

		// Remaining test for success case, so if the case is for fail case, skip it
		if isExpectedError {
			continue
		}

		// When actual value isn't equal expected value
		if accToken == nil {
			t.Fatalf("Case %d: Generated access token is nil", caseNum)
		}
		if refToken == nil {
			t.Fatalf("Case %d: Generated refresh token is nil", caseNum)
		}
		if accToken.AdminUserID != expectedUserID {
			t.Fatalf("Case %d: Not expected user has generated access token.\nExpected:\t%d,\nActual:\t%d", caseNum, expectedUserID, accToken.AdminUserID)
		}
		if refToken.AdminUserID != expectedUserID {
			t.Fatalf("Case %d: Not expected user has generated refresh token.\nExpected:\t%d,\nActual:\t%d", caseNum, expectedUserID, refToken.AdminUserID)
		}
	}
}

func TestAdminAuthModelImpl_RenewAccessToken(t *testing.T) {
	// Initialize model
	aam, aamFin := initAdminAuthModel(t)
	defer aamFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingRefToken string
		ExpectedUserID  uint
	}{
		{IsExpectedError: false, TestingRefToken: "111", ExpectedUserID: 1},
		{IsExpectedError: true, TestingRefToken: "222"},
		{IsExpectedError: true, TestingRefToken: "777"},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		testRefToken := v.TestingRefToken
		expectedUserID := v.ExpectedUserID

		accToken, refToken, err := aam.RenewAccessToken(testRefToken)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When actual value isn't equal expected value
		if isExpectedError && accToken != nil {
			t.Fatalf("Case %d: Generated access token isn't nil", caseNum)
		}
		if isExpectedError && accToken != nil {
			t.Fatalf("Case %d: Generated access token isn't nil", caseNum)
		}

		// Remaining test for success case, so if the case is for fail case, skip it
		if isExpectedError {
			continue
		}

		// When actual value isn't equal expected value
		if accToken == nil {
			t.Fatalf("Case %d: Generated access token is nil", caseNum)
		}
		if refToken == nil {
			t.Fatalf("Case %d: Generated refresh token is nil", caseNum)
		}
		if accToken.AdminUserID != expectedUserID {
			t.Fatalf("Case %d: Not expected user has generated access token.\nExpected:\t%d,\nActual:\t%d", caseNum, expectedUserID, accToken.AdminUserID)
		}
		if refToken.AdminUserID != expectedUserID {
			t.Fatalf("Case %d: Not expected user has generated refresh token.\nExpected:\t%d,\nActual:\t%d", caseNum, expectedUserID, refToken.AdminUserID)
		}
	}
}

type adminModelStub struct{}

func (a *adminModelStub) Add(user *entity.AdminUser) (added *entity.AdminUser, err error) {
	panic("not implemented")
}

func (a *adminModelStub) FindOne(id uint) (user *entity.AdminUser, err error) {
	panic("not implemented")
}

func (a *adminModelStub) FindOneByName(name string) (user *entity.AdminUser, err error) {
	// User 1
	if name == "user_1" {
		return &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "hashed_password_1", Model: gorm.Model{ID: 1}}, nil
	}

	// User 2
	if name == "user_2" {
		return &entity.AdminUser{Name: "user_2", ScreenName: "screen_2", Password: "hashed_password_2", Model: gorm.Model{ID: 2}}, nil
	}

	// Not found
	if name == "error_user" {
		return nil, nil
	}

	return nil, errors.New("adminModelStub.FindOneByName(): Specified name isn't expected as test case")
}

func (a *adminModelStub) Find() (users []entity.AdminUser, err error) {
	panic("not implemented")
}

func (a *adminModelStub) Update(updating *entity.AdminUser) (updated *entity.AdminUser, err error) {
	panic("not implemented")
}

func (a *adminModelStub) Delete(id uint) (err error) {
	panic("not implemented")
}

type hashModelStub struct{}

func (h *hashModelStub) Generate(pass string) (hashed string, err error) {
	if pass == "password_1" {
		return "hashed_password_1", nil
	}

	if pass == "password_2" {
		return "hashed_password_2", nil
	}

	if pass == "password_3" {
		return "hashed_password_3", nil
	}

	if pass == "error_password" {
		return "hashed_error_password", nil
	}

	return "", errors.New("hashModelStub.Generate(): Specified password isn't expected as test case")
}

func (h *hashModelStub) Equals(hashedPassword, rawPassword string) (err error) {
	if hashedPassword == "hashed_password_1" && rawPassword == "password_1" {
		return nil
	}

	if hashedPassword == "hashed_password_2" && rawPassword == "password_2" {
		return nil
	}

	if hashedPassword == "hashed_password_3" && rawPassword == "password_3" {
		return nil

	}

	if hashedPassword == "hashed_error_password" && rawPassword == "error_password" {
		return nil
	}

	return errors.New("hashModelStub.Generate(): Specified password isn't expected as test case")
}
