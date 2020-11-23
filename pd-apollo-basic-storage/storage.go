package main

import (
	"fmt"

	framework "github.com/roonyh/plugin-poc/pd-apollo-framework"
)

type basicStorage struct {
}

func (s basicStorage) Save() {
	fmt.Println("Successfully saved")
}

func GetStorage() framework.Storage {
	storage := basicStorage{}
	return storage
}
