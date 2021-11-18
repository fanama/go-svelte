package database

import (
	"testing"
)

func Test_GetPath(t *testing.T) {
	res, err := GetPath()
	if err != nil {
		t.Errorf("error : %v", err)
	}
	t.Logf(res)
}

func Test_Creation(t *testing.T) {
	err := CreateDB("elements")

	if err != nil {
		t.Errorf("error : %v", err)
	}
}

func Test_Insert(t *testing.T) {
	heroes := map[string]string{"name": "rochi", "power": "500"}

	err := Write("heroes", heroes)

	if err != nil {
		t.Errorf("error : %v", err)
	}

}

// func Test_Delete(t *testing.T) {
// 	heroe := map[string]string{"name": "rochi", "power": "500"}

// 	err := Delete("heroes", heroe)

// 	if err != nil {
// 		t.Errorf("error : %v", err)
// 	}

// }

func Test_Read(t *testing.T) {

	result, err := Read("heroes")

	if err != nil {
		t.Errorf("error : %v", err)
	}

	t.Logf("results : %v", result)

}

func Test_DeleteDB(t *testing.T) {

	err := DeleteDB("heroes")

	if err != nil {
		t.Errorf("error : %v", err)
	}

}
