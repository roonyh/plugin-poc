package main

import (
	"fmt"

	dep "github.com/roonyh/plugin-poc/common-dep"
	framework "github.com/roonyh/plugin-poc/pd-apollo-framework"
)

type basicStorage struct {
}

func (s basicStorage) Save() {
	dep.UseDep()
	fmt.Println("Successfully saved")
}

func GetStorage() framework.Storage {
	storage := basicStorage{}
	return storage
}
