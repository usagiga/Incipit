package model

import (
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/testutils"
	"reflect"
	"testing"
)

func initLinkModel(t *testing.T) (lm LinkModel, finalizer testutils.Finalizer) {
	errChan := make(chan error, 3)
	db, dbFin := testutils.SetupTestDB(t)

	// Insert test values
	go func() {
		errChan <- db.Create(&entity.Link{URL: "https://example.com/1/"}).Error
		errChan <- db.Create(&entity.Link{URL: "https://example.com/2/"}).Error
		errChan <- db.Create(&entity.Link{URL: "https://example.com/3/"}).Error
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			t.Fatal("initLinkModel(): Can't insert testing values; ", err)
		}
	}

	return &LinkModelImpl{db: db, linkValidator: &linkValidatorStub{}}, dbFin
}

func linkEquals(x *entity.Link, y *entity.Link) bool {
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

func TestLinkModelImpl_Add(t *testing.T) {
	// Initialize model
	lm, lmFin := initLinkModel(t)
	defer lmFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingValue    *entity.Link
		ExpectingValue  *entity.Link
	}{
		{
			IsExpectedError: false,
			TestingValue:    &entity.Link{URL: "https://example.com/added/"},
			ExpectingValue:  &entity.Link{URL: "https://example.com/added/", Model: gorm.Model{ID: 4}},
		},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		addingVal := v.TestingValue
		expectedVal := v.ExpectingValue

		actualVal, err := lm.Add(addingVal)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When actual value isn't equal expected value
		if !linkEquals(expectedVal, actualVal) {
			t.Fatalf("Case %d: Actual value isn't equal expected value.\nExpected:\t%v,\nActual:\t%v", caseNum, expectedVal, actualVal)
		}
	}
}

func TestLinkModelImpl_Find(t *testing.T) {
	caseNum := 1

	// Initialize model
	lm, lmFin := initLinkModel(t)
	defer lmFin()

	users, err := lm.Find()

	// When raising NOT expected error
	if err != nil {
		t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
	}

	// When actual value isn't equal expected value
	if len(users) != 3 {
		t.Fatalf("Case %d: The result is too many or too few", caseNum)
	}
}

func TestLinkModelImpl_FindOne(t *testing.T) {
	// Initialize model
	lm, lmFin := initLinkModel(t)
	defer lmFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingID       uint
		ExpectedValue   *entity.Link
	}{
		{
			IsExpectedError: false,
			TestingID:       2,
			ExpectedValue:   &entity.Link{URL: "https://example.com/2/", Model: gorm.Model{ID: 2}},
		},
		{
			IsExpectedError: true,
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

		actualVal, err := lm.FindOne(findingID)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When actual value isn't equal expected value
		if !linkEquals(expectedVal, actualVal) {
			t.Fatalf("Case %d: Actual value isn't equal expected value.\nExpected:\t%v,\nActual:\t%v", caseNum, expectedVal, actualVal)
		}
	}
}

func TestLinkModelImpl_FindOneByShortID(t *testing.T) {
	// Initialize model
	lm, lmFin := initLinkModel(t)
	defer lmFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingShortID  string
		ExpectedValue   *entity.Link
	}{
		{
			IsExpectedError: false,
			TestingShortID:  "2", // 2
			ExpectedValue:   &entity.Link{URL: "https://example.com/2/", Model: gorm.Model{ID: 2}},
		},
		{
			IsExpectedError: true,
			TestingShortID:  "ABCDEFG", // Big integer
			ExpectedValue:   nil,
		},
		{
			IsExpectedError: true,
			TestingShortID:  "+", // Unexpected token
			ExpectedValue:   nil,
		},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		findingShortID := v.TestingShortID
		expectedVal := v.ExpectedValue

		actualVal, err := lm.FindOneByShortID(findingShortID)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When actual value isn't equal expected value
		if !linkEquals(expectedVal, actualVal) {
			t.Fatalf("Case %d: Actual value isn't equal expected value.\nExpected:\t%v,\nActual:\t%v", caseNum, expectedVal, actualVal)
		}
	}
}

func TestLinkModelImpl_Update(t *testing.T) {
	// Initialize model
	lm, lmFin := initLinkModel(t)
	defer lmFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingValue    *entity.Link
		ExpectingValue  *entity.Link
	}{
		{
			IsExpectedError: false,
			TestingValue:    &entity.Link{URL: "https://example.com/updated/", Model: gorm.Model{ID: 2}},
			ExpectingValue:  &entity.Link{URL: "https://example.com/updated/", Model: gorm.Model{ID: 2}},
		},
		{
			IsExpectedError: true,
			TestingValue:    &entity.Link{URL: "https://example.com/1/", Model: gorm.Model{ID: 1}},
			ExpectingValue:  &entity.Link{URL: "https://example.com/1/", Model: gorm.Model{ID: 1}},
		},
		{
			IsExpectedError: true,
			TestingValue:    &entity.Link{Model: gorm.Model{ID: 10}},
			ExpectingValue:  nil,
		},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		updatingVal := v.TestingValue
		expectedVal := v.ExpectingValue

		actualVal, err := lm.Update(updatingVal)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When actual value isn't equal expected value
		if !linkEquals(expectedVal, actualVal) {
			t.Fatalf("Case %d: Actual value isn't equal expected value.\nExpected:\t%v,\nActual:\t%v", caseNum, expectedVal, actualVal)
		}
	}
}

func TestLinkModelImpl_Delete(t *testing.T) {
	// Initialize model
	lm, lmFin := initLinkModel(t)
	defer lmFin()

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

		err := lm.Delete(deletingID)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When not deleting
		deletedVal, _ := lm.FindOne(deletingID)
		if deletedVal != nil {
			t.Fatalf("Case %d: Not deleted", caseNum)
		}
	}
}

type linkValidatorStub struct{}

func (v *linkValidatorStub) ValidateAll(link *entity.Link) (err error) {
	return nil
}

func (v *linkValidatorStub) ValidateID(id uint) (err error) {
	return nil
}

func (v *linkValidatorStub) ValidateURL(url string) (err error) {
	return nil
}
