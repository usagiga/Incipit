package entity

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
	"kkn.fi/base62"
)

// Link represents URL to shorten
type Link struct {
	gorm.Model
	URL string
}

// GetShortID gets ID as base62 string from its actual ID
func (e *Link) GetShortID() string {
	return base62.Encode(int64(e.ID))
}

// ToActualID gets actual ID from specified short ID
func ToActualID(shortId string) (actualId uint, err error) {
	rawId, err := base62.Decode(shortId)
	if err != nil {
		return 0, xerrors.Errorf("Can't parse actual id: %w", err)
	}

	return uint(rawId), nil
}
