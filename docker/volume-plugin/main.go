package main

import (
	"fmt"

	"github.com/docker/go-plugins-helpers/volume"
)

type driver struct{}

func main() {
	h := volume.NewHandler(driver{})
	if err := h.ServeUnix("root", "test_volume"); err != nil {
		panic(err.Error())
	}
}

func (driver) Create(volume.Request) volume.Response {
	fmt.Println("Fake create now")
	return volume.Response{}
}

func (driver) Remove(volume.Request) volume.Response {
	fmt.Println("Fake remove now")
	return volume.Response{}
}

func (driver) Path(volume.Request) volume.Response {
	fmt.Println("Fake path now")
	return volume.Response{}
}

func (driver) Mount(volume.Request) volume.Response {
	fmt.Println("Fake mount now")
	return volume.Response{}
}

func (driver) Unmount(volume.Request) volume.Response {
	fmt.Println("Fake rem now")
	return volume.Response{}
}

func (driver) Get(volume.Request) volume.Response {
	fmt.Println("Fake get  now")
	return volume.Response{Err: "No such volume"}
}

func (driver) List(volume.Request) volume.Response {
	fmt.Println("Fake list now")
	return volume.Response{}
}
