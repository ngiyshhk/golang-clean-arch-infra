package database

import (
	"testing"

	"github.com/ngiyshhk/golang-clean-arch-usecase/repository"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

func truncate(t *testing.T, repo repository.UserRepository) {
	users, err := repo.Select([]int{})
	if err != nil {
		t.Fatal("failed to select all")
	}

	for _, user := range users {
		repo.Delete(user.ID)
	}
}

func TestCreateNormal(t *testing.T) {
	repo := &UserRepositoryImpl{}
	truncate(t, repo)

	user := &model.User{ID: 1, Name: "user1", Age: 20}
	res, err := repo.Insert(user)
	if res != true {
		t.Errorf("res is not true. got=%t", res)
	}
	if err != nil {
		t.Errorf("err is not nil. got=%v", err)
	}
}

func TestCreateError(t *testing.T) {
	repo := &UserRepositoryImpl{}
	truncate(t, repo)

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
	repo := &UserRepositoryImpl{}
	truncate(t, repo)

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
	repo := &UserRepositoryImpl{}
	truncate(t, repo)

	user := &model.User{ID: 1, Name: "user1", Age: 20}
	res, err := repo.Update(user)
	if res != false {
		t.Errorf("res is not false. got=%t", res)
	}
	if err == nil {
		t.Error("err is nil")
	}
}

func TestDeleteNormal(t *testing.T) {
	repo := &UserRepositoryImpl{}
	truncate(t, repo)

	res, err := repo.Delete(1)
	if res != true {
		t.Errorf("res is not true. got=%t", res)
	}
	if err != nil {
		t.Errorf("err is not nil. got=%v", err)
	}
}

func TestDeleteError(t *testing.T) {
	repo := &UserRepositoryImpl{}
	truncate(t, repo)

	res, err := repo.Delete(1)
	if res != false {
		t.Errorf("res is not false. got=%t", res)
	}
	if err == nil {
		t.Error("err is nil")
	}
}

func TestSelectNormal(t *testing.T) {
	user1 := &model.User{ID: 1, Name: "user1", Age: 20}
	user2 := &model.User{ID: 2, Name: "user2", Age: 21}
	userFixtures := []*model.User{user1, user2}

	repo := &UserRepositoryImpl{}
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

func TestSelectError(t *testing.T) {
	repo := &UserRepositoryImpl{}
	res, err := repo.Select([]int{1, 2})
	if res != nil {
		t.Errorf("res is not nil. got=%v", res)
	}
	if err == nil {
		t.Error("err is nil")
	}
}
