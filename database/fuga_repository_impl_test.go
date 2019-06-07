package database

import (
	"testing"

	"github.com/ngiyshhk/golang-clean-arch-usecase/repository"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

func truncate(t *testing.T, repo repository.FugaRepository) {
	fugas, err := repo.Select([]int{})
	if err != nil {
		t.Fatal("failed to select all")
	}

	for _, fuga := range fugas {
		repo.Delete(fuga.ID)
	}
}

func TestCreateNormal(t *testing.T) {
	repo := &FugaRepositoryImpl{}
	truncate(t, repo)

	fuga := &model.Fuga{ID: 1, Name: "fuga1", Age: 20}
	res, err := repo.Insert(fuga)
	if res != true {
		t.Errorf("res is not true. got=%t", res)
	}
	if err != nil {
		t.Errorf("err is not nil. got=%v", err)
	}
}

func TestCreateError(t *testing.T) {
	repo := &FugaRepositoryImpl{}
	truncate(t, repo)

	fuga := &model.Fuga{ID: 1, Name: "fuga1", Age: 20}
	res, err := repo.Insert(fuga)
	if res != false {
		t.Errorf("res is not false. got=%t", res)
	}
	if err == nil {
		t.Error("err is nil")
	}
}

func TestUpdateNormal(t *testing.T) {
	repo := &FugaRepositoryImpl{}
	truncate(t, repo)

	fuga := &model.Fuga{ID: 1, Name: "fuga1", Age: 20}
	res, err := repo.Update(fuga)
	if res != true {
		t.Errorf("res is not true. got=%t", res)
	}
	if err != nil {
		t.Errorf("err is not nil. got=%v", err)
	}
}

func TestUpdateError(t *testing.T) {
	repo := &FugaRepositoryImpl{}
	truncate(t, repo)

	fuga := &model.Fuga{ID: 1, Name: "fuga1", Age: 20}
	res, err := repo.Update(fuga)
	if res != false {
		t.Errorf("res is not false. got=%t", res)
	}
	if err == nil {
		t.Error("err is nil")
	}
}

func TestDeleteNormal(t *testing.T) {
	repo := &FugaRepositoryImpl{}
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
	repo := &FugaRepositoryImpl{}
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
	fuga1 := &model.Fuga{ID: 1, Name: "fuga1", Age: 20}
	fuga2 := &model.Fuga{ID: 2, Name: "fuga2", Age: 21}
	fugaFixtures := []*model.Fuga{fuga1, fuga2}

	repo := &FugaRepositoryImpl{}
	res, err := repo.Select([]int{1, 2})
	if len(res) != len(fugaFixtures) {
		t.Errorf("len(res) is not %d. got=%d", len(fugaFixtures), len(res))
	}
	for i := 0; i < len(res); i++ {
		if res[i].ID != fugaFixtures[i].ID {
			t.Errorf("res[%d].ID is not %d. got=%d", i, fugaFixtures[i].ID, res[i].ID)
			t.Errorf("res[%d].Name is not %s. got=%s", i, fugaFixtures[i].Name, res[i].Name)
			t.Errorf("res[%d].Age is not %d. got=%d", i, fugaFixtures[i].Age, res[i].Age)
		}
	}
	if err != nil {
		t.Errorf("err is not nil. got=%v", err)
	}
}

func TestSelectError(t *testing.T) {
	repo := &FugaRepositoryImpl{}
	res, err := repo.Select([]int{1, 2})
	if res != nil {
		t.Errorf("res is not nil. got=%v", res)
	}
	if err == nil {
		t.Error("err is nil")
	}
}
