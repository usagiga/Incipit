package model

import (
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
	interr "github.com/usagiga/Incipit/back/entity/errors"
)

type LinkModelImpl struct {
	db *gorm.DB
}

func NewLinkModel(db *gorm.DB) LinkModel {
	return &LinkModelImpl{db: db}
}

func (m *LinkModelImpl) Add(link *entity.Link) (added *entity.Link, err error) {
	// Add the row
	result := m.db.Create(link)
	err = result.Error
	if err != nil {
		return nil, interr.NewDistinctError("Can't add link", interr.LinkModel, interr.LinkModel_FailedAdd, nil).Wrap(err)
	}
	if result.RowsAffected <= 0 {
		return nil, interr.NewDistinctError("Can't add link", interr.LinkModel, interr.LinkModel_FailedAdd, nil)
	}

	return link, nil
}

func (m *LinkModelImpl) FindOne(id uint) (link *entity.Link, err error) {
	link = &entity.Link{}

	// Find the row
	result := m.db.First(link, id)
	err = result.Error
	if result.RecordNotFound() {
		return nil, interr.NewDistinctError("Link isn't found", interr.LinkModel, interr.LinkModel_FindingLinkNotFound, nil).Wrap(err)
	}
	if err != nil {
		return nil, interr.NewDistinctError("Can't find link", interr.LinkModel, interr.LinkModel_FailedFind, nil).Wrap(err)
	}

	return link, nil
}

func (m *LinkModelImpl) Find() (links []entity.Link, err error) {
	// Find all rows
	result := m.db.Find(&links)
	err = result.Error
	if result.RecordNotFound() {
		return nil, interr.NewDistinctError("Link aren't found", interr.LinkModel, interr.LinkModel_FindingLinkNotFound, nil).Wrap(err)
	}
	if err != nil {
		return nil, interr.NewDistinctError("Can't find link", interr.LinkModel, interr.LinkModel_FailedFind, nil).Wrap(err)
	}

	return links, nil
}

func (m *LinkModelImpl) Update(updating *entity.Link) (updated *entity.Link, err error) {
	// Finding the row
	_, err = m.FindOne(updating.ID)
	if err != nil {
		return nil, interr.NewDistinctError("Can't find link", interr.LinkModel, interr.LinkModel_UpdatingLinkNotFound, nil).Wrap(err)
	}

	// Update the row
	result := m.db.Model(&entity.Link{}).Update(updating)
	err = result.Error
	if err != nil {
		return nil, interr.NewDistinctError("Can't update link", interr.LinkModel, interr.LinkModel_FailedUpdate, nil).Wrap(err)
	}
	if result.RowsAffected <= 0 {
		// There's no value to change
		return updating, nil
	}

	return updating, nil
}

func (m *LinkModelImpl) Delete(id uint) (err error) {
	// Delete the row
	result := m.db.Delete(&entity.Link{}, id)
	err = result.Error
	if err != nil {
		return interr.NewDistinctError("Can't delete link", interr.LinkModel, interr.LinkModel_FailedDelete, nil).Wrap(err)
	}
	if result.RowsAffected <= 0 {
		return interr.NewDistinctError("Can't delete link", interr.LinkModel, interr.LinkModel_FailedDelete, nil)
	}

	return nil
}
