package main

import (
	framework "github.com/roonyh/plugin-poc/pd-apollo-framework"
)

func main() {
	f := framework.Init()
	f.Storage.Save()
}
