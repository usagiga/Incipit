package model

import "golang.org/x/crypto/bcrypt"

type HashModelImpl struct {}

func NewHashModel() HashModel {
	return &HashModelImpl{}
}

func (h *HashModelImpl) Generate(pass string) (hashed string, err error) {
	passBytes := []byte(pass)

	hashedBytes, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashed = string(hashedBytes)

	return hashed, nil
}

func (h *HashModelImpl) Equals(hashedPassword, rawPassword string) (err error) {
	hashedBytes := []byte(hashedPassword)
	rawBytes := []byte(rawPassword)

	return bcrypt.CompareHashAndPassword(hashedBytes, rawBytes)
}

