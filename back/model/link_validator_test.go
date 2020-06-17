package model

import (
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/testutils"
	"testing"
)

func initLinkValidator() (lv LinkValidator, finalizer testutils.Finalizer) {
	lv = NewLinkValidator("incipit.com")

	finalizer = func() {}

	return lv, finalizer
}

func TestLinkValidatorImpl_ValidateAll(t *testing.T) {
	// Initialize model
	lv, lvFin := initLinkValidator()
	defer lvFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingLink     *entity.Link
	}{
		{IsExpectedError: false, TestingLink: &entity.Link{URL: "http://example.com/1/"}},
		{IsExpectedError: false, TestingLink: &entity.Link{URL: "https://example.com/1/"}},
		{IsExpectedError: false, TestingLink: &entity.Link{URL: "ftp://example.com/1/"}},
		{IsExpectedError: true, TestingLink: &entity.Link{URL: "UNKNOWN://example.com/1/"}},
		{IsExpectedError: true, TestingLink: &entity.Link{URL: "UNKNOWN SCHEME"}},
		{IsExpectedError: true, TestingLink: &entity.Link{URL: ""}},
		{IsExpectedError: true, TestingLink: &entity.Link{URL: "https://incipit.com/"}},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		testingLink := v.TestingLink

		err := lv.ValidateAll(testingLink)

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

func TestLinkValidatorImpl_ValidateID(t *testing.T) {
	// NOP
}

func TestLinkValidatorImpl_ValidateURL(t *testing.T) {
	// Initialize model
	lv, lvFin := initLinkValidator()
	defer lvFin()

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingUrl      string
	}{
		{IsExpectedError: false, TestingUrl: "http://example.com/1/"},
		{IsExpectedError: false, TestingUrl: "https://example.com/1/"},
		{IsExpectedError: false, TestingUrl: "ftp://example.com/1/"},
		{IsExpectedError: true, TestingUrl: "UNKNOWN://example.com/1/"},
		{IsExpectedError: true, TestingUrl: "UNKNOWN"},
		{IsExpectedError: true, TestingUrl: ""},
		{IsExpectedError: true, TestingUrl: "https://incipit.com/"},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		testingLink := v.TestingUrl

		err := lv.ValidateURL(testingLink)

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
