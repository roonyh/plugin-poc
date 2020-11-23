package framework

import (
	"fmt"
	"os"
	"plugin"
)

type Storage interface {
	Save()
}

func Init() {
	fmt.Println("Apollo framwork 0.0.1")
	storagePlug, err := plugin.Open("./storage.so")
	if err != nil {
		fmt.Println("Storage plugin not found")
		fmt.Println(err)
		os.Exit(1)
	}

	storageFactorySymbol, err := storagePlug.Lookup("GetStorage")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	storageFactory, ok := storageFactorySymbol.(func() Storage)
	if !ok {
		fmt.Println("Storage interface is not implemented properly")
		os.Exit(1)
	}
	storage := storageFactory()
	storage.Save()
}
