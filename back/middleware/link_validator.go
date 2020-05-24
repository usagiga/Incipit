package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity"
)

type LinkValidatorImpl struct {

}

func NewLinkValidator() LinkValidator {
	return &LinkValidatorImpl{}
}

func (m *LinkValidatorImpl) Handle(c *gin.Context) {
	panic("implement me")
}

func (m *LinkValidatorImpl) ValidateAll(link *entity.Link) (err error) {
	panic("implement me")
}

func (m *LinkValidatorImpl) ValidateID(id uint) (err error) {
	panic("implement me")
}

func (m *LinkValidatorImpl) ValidateURL(url string) (err error) {
	panic("implement me")
}
