package database

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"
	"github.com/ngiyshhk/golang-clean-arch-usecase/repository"
)

var user1 = &model.User{ID: 1, Name: "user1", Age: 20}
var user2 = &model.User{ID: 2, Name: "user2", Age: 21}
var userFixtures = []*model.User{user1, user2}

func truncate(t *testing.T, repo repository.UserRepository) {
	users, err := repo.Select([]int{})
	if err != nil {
		t.Fatal("failed to select all")
	}

	for _, user := range users {
		repo.Delete(user.ID)
	}
}

func fixtures(t *testing.T, repo repository.UserRepository) {
	for _, user := range userFixtures {
		_, err := repo.Insert(user)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCreateNormal(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@/clean_arch?charset=utf8&parseTime=True")
	if err != nil {
		t.Errorf("db connection error. got=%v", err)
	}

	repo := &UserRepositoryImpl{Db: db}
	truncate(t, repo)
	fixtures(t, repo)

	user := &model.User{ID: 11, Name: "user11", Age: 30}
	res, err := repo.Insert(user)
	if res != true {
		t.Errorf("res is not true. got=%t", res)
	}
	if err != nil {
		t.Errorf("err is not nil. got=%v", err)
	}
}

func TestCreateError(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@/clean_arch?charset=utf8&parseTime=True")
	if err != nil {
		t.Errorf("db connection error. got=%v", err)
	}

	repo := &UserRepositoryImpl{Db: db}
	truncate(t, repo)
	fixtures(t, repo)

	user := &model.User{ID: 1, Name: "user1", Age: 20}
	res, err := repo.Insert(user)
	if res != false {
		t.Errorf("res is not false. got=%t", res)
	}
	if err == nil {
		t.Error("err is nil")
	}
}

func TestUpdateNormal(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@/clean_arch?charset=utf8&parseTime=True")
	if err != nil {
		t.Errorf("db connection error. got=%v", err)
	}

	repo := &UserRepositoryImpl{Db: db}
	truncate(t, repo)
	fixtures(t, repo)

	user := &model.User{ID: 1, Name: "user1", Age: 20}
	res, err := repo.Update(user)
	if res != true {
		t.Errorf("res is not true. got=%t", res)
	}
	if err != nil {
		t.Errorf("err is not nil. got=%v", err)
	}
}

func TestUpdateError(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@/clean_arch?charset=utf8&parseTime=True")
	if err != nil {
		t.Errorf("db connection error. got=%v", err)
	}

	repo := &UserRepositoryImpl{Db: db}
	truncate(t, repo)
	fixtures(t, repo)

	user := &model.User{ID: 11, Name: "user1", Age: 20}
	res, err := repo.Update(user)
	if res != false {
		t.Errorf("res is not false. got=%t", res)
	}
	if err == nil {
		t.Error("err is nil")
	}
}

func TestDeleteNormal(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@/clean_arch?charset=utf8&parseTime=True")
	if err != nil {
		t.Errorf("db connection error. got=%v", err)
	}

	repo := &UserRepositoryImpl{Db: db}
	truncate(t, repo)
	fixtures(t, repo)

	res, err := repo.Delete(1)
	if res != true {
		t.Errorf("res is not true. got=%t", res)
	}
	if err != nil {
		t.Errorf("err is not nil. got=%v", err)
	}
}

func TestSelectNormal(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@/clean_arch?charset=utf8&parseTime=True")
	if err != nil {
		t.Errorf("db connection error. got=%v", err)
	}

	repo := &UserRepositoryImpl{Db: db}
	truncate(t, repo)
	fixtures(t, repo)

	res, err := repo.Select([]int{1, 2})
	if len(res) != len(userFixtures) {
		t.Errorf("len(res) is not %d. got=%d", len(userFixtures), len(res))
	}
	for i := 0; i < len(res); i++ {
		if res[i].ID != userFixtures[i].ID {
			t.Errorf("res[%d].ID is not %d. got=%d", i, userFixtures[i].ID, res[i].ID)
			t.Errorf("res[%d].Name is not %s. got=%s", i, userFixtures[i].Name, res[i].Name)
			t.Errorf("res[%d].Age is not %d. got=%d", i, userFixtures[i].Age, res[i].Age)
		}
	}
	if err != nil {
		t.Errorf("err is not nil. got=%v", err)
	}
}
