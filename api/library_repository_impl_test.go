package api

import (
	"fmt"
	"testing"
)

func TestGetAllNormal(t *testing.T) {
	repo := &LibraryRepositoryImpl{}

	res, err := repo.GetAll()
	if err != nil {
		t.Fatalf("err is not nil. got=%v", err)
	}
	if len(*res) == 0 {
		t.Errorf("len(res) is empty.")
	}
	for _, library := range *res {
		fmt.Println(*library)
	}
}
