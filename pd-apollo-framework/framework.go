package framework

import (
	"fmt"
	"os"
	"plugin"
)

type Storage interface {
	Save()
}

type Receiver interface {
	Start()
}

type Framework struct {
	Storage  Storage
	Receiver Receiver
}

func Init() *Framework {
	fmt.Println("Apollo framwork 0.0.1")
	storagePlug, err := plugin.Open("./storage.so")
	if err != nil {
		fmt.Println("Storage plugin not found")
		fmt.Println(err)
		os.Exit(1)
	}

	receiverPlug, err := plugin.Open("./receiver.so")
	if err != nil {
		fmt.Println("Receiver plugin not found")
		fmt.Println(err)
		os.Exit(1)
	}

	storageFactorySymbol, err := storagePlug.Lookup("GetStorage")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	receiverFactorySymbol, err := receiverPlug.Lookup("GetReceiver")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	storageFactory, ok := storageFactorySymbol.(func() Storage)
	if !ok {
		fmt.Println("Storage interface is not implemented properly")
		os.Exit(1)
	}

	receiverFactory, ok := receiverFactorySymbol.(func() Receiver)
	if !ok {
		fmt.Println("Storage interface is not implemented properly")
		os.Exit(1)
	}

	storage := storageFactory()
	receiver := receiverFactory()
	framework := Framework{Storage: storage, Receiver: receiver}

	return &framework
}
