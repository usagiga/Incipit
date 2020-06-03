package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
)

type LinkModelImpl struct {
	db *gorm.DB
}

func NewLinkModel() LinkModel {
	return &LinkModelImpl{}
}

func (m *LinkModelImpl) Add(link *entity.Link) (added *entity.Link, err error) {
	// Add the row
	result := m.db.Create(link)
	err = result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, errors.New("LinkModel.Add(): Can't add it")
	}

	return link, err
}

func (m *LinkModelImpl) FindOne(id uint) (link *entity.Link, err error) {
	link = &entity.Link{}

	// Find the row
	result := m.db.First(link, id)
	err = result.Error
	if result.RecordNotFound() {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return link, err
}

func (m *LinkModelImpl) Find() (links []entity.Link, err error) {
	// Find all rows
	result := m.db.Find(&links)
	err = result.Error
	if result.RecordNotFound() {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return links, err
}

func (m *LinkModelImpl) Update(updating *entity.Link) (updated *entity.Link, err error) {
	// Finding the row
	found, err := m.FindOne(updating.ID)
	if err != nil {
		return nil, err
	}
	if found == nil {
		return nil, errors.New("LinkModel.Update(): Not found updating row")
	}

	// Update the row
	result := m.db.Model(&entity.Link{}).Update(updating)
	err = result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return updating, nil
	}

	return updating, err
}

func (m *LinkModelImpl) Delete(id uint) (err error) {
	// Delete the row
	result := m.db.Delete(&entity.Link{}, id)
	err = result.Error
	if err != nil {
		return nil
	}
	if result.RowsAffected <= 0 {
		return errors.New("LinkModel.Delete(): Can't delete it")
	}

	return nil
}
