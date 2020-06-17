package model

import (
	"github.com/usagiga/Incipit/back/entity"
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"golang.org/x/xerrors"
	urllib "net/url"
)

type LinkValidatorImpl struct {
	incipitHost string
}

func NewLinkValidator(incipitHost string) LinkValidator {
	return &LinkValidatorImpl{incipitHost: incipitHost}
}

func (m *LinkValidatorImpl) ValidateAll(link *entity.Link) (err error) {
	errChan := make(chan error, 2)

	go func() {
		errChan <- m.ValidateID(link.ID)
		errChan <- m.ValidateURL(link.URL)
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return xerrors.Errorf("This link is invalid: %w", err)
		}
	}

	return nil
}

func (m *LinkValidatorImpl) ValidateID(id uint) (err error) {
	return nil
}

func (m *LinkValidatorImpl) ValidateURL(url string) (err error) {
	errChan := make(chan error, 2)
	incipitHost := m.incipitHost

	go func() {
		parsed, err := urllib.ParseRequestURI(url)
		if err != nil {
			errChan <- interr.NewDistinctError("Invalid URL", interr.LinkValidation, interr.LinkValidation_URLIsInvalid, nil).Wrap(err)
		}
		if parsed.Host == incipitHost {
			errChan <- interr.NewDistinctError("This URL is itself", interr.LinkValidation, interr.LinkValidation_URLIsIncipit, nil)
		}
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return xerrors.Errorf("This URL is invalid: %w", err)
		}
	}

	return nil
}
