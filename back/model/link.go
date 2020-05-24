package model

import (
	"github.com/usagiga/Incipit/back/entity"
)

type LinkModelImpl struct {

}

func NewLinkModel() LinkModel {
	return &LinkModelImpl{}
}

func (m *LinkModelImpl) Add(user *entity.Link) (added *entity.Link, err error) {
	panic("implement me")
}

func (m *LinkModelImpl) FindOne(id uint) (link *entity.Link, err error) {
	panic("implement me")
}

func (m *LinkModelImpl) Find() (links []entity.Link, err error) {
	panic("implement me")
}

func (m *LinkModelImpl) Update(updating *entity.Link) (updated *entity.Link, err error) {
	panic("implement me")
}

func (m *LinkModelImpl) Delete(id uint) (err error) {
	panic("implement me")
}
