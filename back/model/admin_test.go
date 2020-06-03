package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/testutils"
	"reflect"
	"testing"
)

func initAdminModel(t *testing.T) (am AdminModel, finalizer testutils.Finalizer) {
	errChan := make(chan error, 3)
	db, dbFin := testutils.SetupTestDB(t)

	// Insert test values
	go func() {
		errChan <- db.Create(&entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "password_1"}).Error
		errChan <- db.Create(&entity.AdminUser{Name: "user_2", ScreenName: "screen_2", Password: "password_2"}).Error
		errChan <- db.Create(&entity.AdminUser{Name: "user_3", ScreenName: "screen_3", Password: "password_3"}).Error
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			t.Fatal("initAdminModel(): Can't insert testing values; ", err)
		}
	}

	return &AdminModelImpl{db: db, hashModel: &hashModelStub{}}, dbFin
}

func adminUserEquals(x *entity.AdminUser, y *entity.AdminUser) bool {
	// If null
	if x == nil || y == nil {
		return x == y
	}

	// Compare ID
	if x.ID != y.ID {
		return false
	}

	// Compare other value without `gorm.Model`
	emptyModel := gorm.Model{}
	x.Model = emptyModel
	y.Model = emptyModel
	if !reflect.DeepEqual(x, y) {
		return false
	}

	return true
}

func TestAdminModelImpl_Add(t *testing.T) {
	// Initialize model
	am, amFin := initAdminModel(t)
	defer amFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingValue    *entity.AdminUser
		ExpectingValue  *entity.AdminUser
	}{
		{
			IsExpectedError: false,
			TestingValue:    &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "password_1"},
			ExpectingValue:  &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "hashed_password_1", Model: gorm.Model{ID: 4}},
		},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		addingVal := v.TestingValue
		expectedVal := v.ExpectingValue

		actualVal, err := am.Add(addingVal)

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

func TestAdminModelImpl_Find(t *testing.T) {
	caseNum := 1

	// Initialize model
	am, amFin := initAdminModel(t)
	defer amFin()

	users, err := am.Find()

	// When raising NOT expected error
	if err != nil {
		t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
	}

	// When actual value isn't equal expected value
	if len(users) != 3 {
		t.Fatalf("Case %d: The result is too many or too few", caseNum)
	}
}

func TestAdminModelImpl_FindOne(t *testing.T) {
	// Initialize model
	am, amFin := initAdminModel(t)
	defer amFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingID       uint
		ExpectedValue   *entity.AdminUser
	}{
		{
			IsExpectedError: false,
			TestingID:       2,
			ExpectedValue:   &entity.AdminUser{Name: "user_2", ScreenName: "screen_2", Password: "password_2", Model: gorm.Model{ID: 2}},
		},
		{
			IsExpectedError: false,
			TestingID:       10,
			ExpectedValue:   nil,
		},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		findingID := v.TestingID
		expectedVal := v.ExpectedValue

		actualVal, err := am.FindOne(findingID)

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

func TestAdminModelImpl_FindOneByName(t *testing.T) {
	// Initialize model
	am, amFin := initAdminModel(t)
	defer amFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingName     string
		ExpectedValue   *entity.AdminUser
	}{
		{
			IsExpectedError: false,
			TestingName:     "user_2",
			ExpectedValue:   &entity.AdminUser{Name: "user_2", ScreenName: "screen_2", Password: "password_2", Model: gorm.Model{ID: 2}},
		},
		{
			IsExpectedError: false,
			TestingName:     "not_existed_user",
			ExpectedValue:   nil,
		},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		findingName := v.TestingName
		expectedVal := v.ExpectedValue

		actualVal, err := am.FindOneByName(findingName)

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

func TestAdminModelImpl_Update(t *testing.T) {
	// Initialize model
	am, amFin := initAdminModel(t)
	defer amFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingValue    *entity.AdminUser
		ExpectingValue  *entity.AdminUser
	}{
		{
			IsExpectedError: false,
			TestingValue:    &entity.AdminUser{Name: "updated_name", ScreenName: "updated_screen_name", Password: "password_3", Model: gorm.Model{ID: 2}},
			ExpectingValue:  &entity.AdminUser{Name: "updated_name", ScreenName: "updated_screen_name", Password: "hashed_password_3", Model: gorm.Model{ID: 2}},
		},
		{
			IsExpectedError: false,
			TestingValue:    &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "password_1", Model: gorm.Model{ID: 1}},
			ExpectingValue:  &entity.AdminUser{Name: "user_1", ScreenName: "screen_1", Password: "hashed_password_1", Model: gorm.Model{ID: 1}},
		},
		{
			IsExpectedError: true,
			TestingValue:    &entity.AdminUser{Model: gorm.Model{ID: 10}},
			ExpectingValue:  nil,
		},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		updatingVal := v.TestingValue
		expectedVal := v.ExpectingValue

		actualVal, err := am.Update(updatingVal)

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

func TestAdminModelImpl_Delete(t *testing.T) {
	// Initialize model
	am, amFin := initAdminModel(t)
	defer amFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingID       uint
	}{
		{
			IsExpectedError: false,
			TestingID:       2,
		},
		{
			IsExpectedError: true,
			TestingID:       10,
		},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		deletingID := v.TestingID

		err := am.Delete(deletingID)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When not deleting
		deletedVal, _ := am.FindOne(deletingID)
		if deletedVal != nil {
			t.Fatalf("Case %d: Not deleted", caseNum)
		}
	}
}
