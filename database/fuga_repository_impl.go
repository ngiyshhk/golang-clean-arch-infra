package database

import (
	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

type FugaRepositoryImpl struct{}

func (_ *FugaRepositoryImpl) Insert(entity *model.Fuga) (bool, error) {
	return false, nil
}

func (_ *FugaRepositoryImpl) Update(entity *model.Fuga) (bool, error) {
	return false, nil
}

func (_ *FugaRepositoryImpl) Select(ids []int) ([]*model.Fuga, error) {
	return make([]*model.Fuga, 0), nil
}

func (_ *FugaRepositoryImpl) Delete(id int) (bool, error) {
	return false, nil
}
