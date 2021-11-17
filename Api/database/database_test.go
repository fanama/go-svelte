package database

import (
	"fmt"
	"testing"
)

func Test_GetPath(t *testing.T) {
	res, err := GetPath()
	if err != nil {
		t.Errorf("error : %v", err)
	}
	fmt.Println(res)
}

func Test_Creation(t *testing.T) {
	err := CreateDB("elements")

	if err != nil {
		t.Errorf("error : %v", err)
	}
}

func Test_Insert(t *testing.T) {
	heroes := map[string]string{"name": "vegeta", "power": "16000"}

	err := Write("elements", heroes)

	if err != nil {
		t.Errorf("error : %v", err)
	}
}
