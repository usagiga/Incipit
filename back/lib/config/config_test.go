package config

import (
	"os"
	"reflect"
	"testing"
)

func initConfig(t *testing.T) {
	errChan := make(chan error, 3)

	go func() {
		errChan <- os.Setenv("INCIPIT_TEST_INT", "12345")
		errChan <- os.Setenv("INCIPIT_TEST_STRING", "string")
		errChan <- os.Setenv("INCIPIT_TEST_BOOL", "false")

		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			t.Fatal("Can't set testing environment variables; ", err)
		}
	}
}

type testConfig1 struct {
	IntVal    int    `env:"INCIPIT_TEST_INT"`
	StringVal string `env:"INCIPIT_TEST_STRING"`
	BoolVal   bool   `env:"INCIPIT_TEST_BOOL"`
}

type testConfig2 struct {
	StringVal string `env:"INCIPIT_TEST_STRING"`
	ErrorVal  string `env:"INCIPIT_TEST_ERROR"`
}

func TestLoad(t *testing.T) {
	// Initialize config
	initConfig(t)

	// Declare test cases
	testCases := []struct {
		IsExpectedError bool
		TestingValue    interface{}
		ExpectingValue  interface{}
	}{
		{IsExpectedError: false, TestingValue: &testConfig1{}, ExpectingValue: &testConfig1{IntVal: 12345, StringVal: "string", BoolVal: false}},
		{IsExpectedError: true, TestingValue: &testConfig2{}},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		isExpectedError := v.IsExpectedError
		actual := v.TestingValue
		expected := v.ExpectingValue

		err := Load(actual)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Fatalf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Fatalf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// Following test is not for error
		if isExpectedError {
			continue
		}

		// When actual value isn't equal expected value
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("Case %d: Actual value isn't equal expected value.\nExpected:\t%v,\nActual:\t%v", caseNum, expected, actual)
		}
	}
}
