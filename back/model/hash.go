package model

import (
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"golang.org/x/crypto/bcrypt"
)

type HashModelImpl struct{}

func NewHashModel() HashModel {
	return &HashModelImpl{}
}

func (h *HashModelImpl) Generate(pass string) (hashed string, err error) {
	passBytes := []byte(pass)

	hashedBytes, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", interr.NewDistinctError("Can't generate hash", interr.HashModel, interr.HashModel_FailedGenerate, nil).Wrap(err)
	}

	hashed = string(hashedBytes)

	return hashed, nil
}

func (h *HashModelImpl) Equals(hashedPassword, rawPassword string) (err error) {
	hashedBytes := []byte(hashedPassword)
	rawBytes := []byte(rawPassword)

	err = bcrypt.CompareHashAndPassword(hashedBytes, rawBytes)
	if err != nil {
		return interr.NewDistinctError("Don't match hash and password", interr.HashModel, interr.HashModel_FailedCompare, nil).Wrap(err)
	}

	return nil
}
