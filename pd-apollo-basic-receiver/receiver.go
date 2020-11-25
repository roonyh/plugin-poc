package main

import (
	"fmt"

	framework "github.com/roonyh/plugin-poc/pd-apollo-framework"

	dep "github.com/roonyh/plugin-poc/common-dep"
)

type receiver struct {
}

func (s receiver) Start() {
	dep.UseDep()
	fmt.Println("Successfully saved")
}

func GetReceiver() framework.Receiver {
	r := &receiver{}
	return r
}
