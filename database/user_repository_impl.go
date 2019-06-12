package database

import (
	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

// UserRepositoryImpl is implement user repository
type UserRepositoryImpl struct{}

// Insert is
func (*UserRepositoryImpl) Insert(entity *model.User) (bool, error) {
	return false, nil
}

// Update is
func (*UserRepositoryImpl) Update(entity *model.User) (bool, error) {
	return false, nil
}

// Select is
func (*UserRepositoryImpl) Select(ids []int) ([]*model.User, error) {
	return make([]*model.User, 0), nil
}

// Delete is
func (*UserRepositoryImpl) Delete(id int) (bool, error) {
	return false, nil
}
