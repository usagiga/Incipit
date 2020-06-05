package main

import (
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/config"
)

func LoadConfig() (result *entity.Config, err error) {
	result = &entity.Config{}

	err = config.Load(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
