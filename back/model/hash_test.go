package model

import "testing"

func initHashModel() HashModel {
	return NewHashModel()
}

func TestHashModel(t *testing.T) {
	caseNum := 1

	// Initialize HashModel
	hm := initHashModel()

	// Do test
	testingPass := "password"
	hashedPass, err := hm.Generate(testingPass)
	if err != nil {
		t.Fatalf("Case %d: Can't generate hash; %v", caseNum, err)
	}

	err = hm.Equals(hashedPass, testingPass)
	if err != nil {
		t.Fatalf("Case %d: Didn't match hash value and real value; %v", caseNum, err)
	}
}
