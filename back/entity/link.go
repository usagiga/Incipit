package entity

import "github.com/jinzhu/gorm"

// Link represents URL to shorten
type Link struct {
	gorm.Model
	URL string
}

// GetShortID gets ID as base62 string from its actual ID
func (e *Link) GetShortID() string {
	panic("not implemented")
}

// ToActualID gets actual ID from specified short ID
func ToActualID(shortId string) uint {
	panic("not implemented")
}
