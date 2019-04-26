package database

import (
	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

type FugaRepositoryImpl struct{}

func (_ *FugaRepositoryImpl) Insert(entity *model.Fuga) (bool, error) {
	return false, nil
}
