package main

import (
	"fmt"
)

type Car interface {
	Drive()
}

type BMW struct{}

func (car BMW) Drive() {
	fmt.Println("BMW is driving ... ")
}

type Mercedes struct{}

func (car Mercedes) Drive() {
	fmt.Println("Mercedes is driving ... ")
}

type CarFactory interface {
	CreateCar() Car
}

type BMWFactory struct{}

func (factory BMWFactory) CreateCar() Car {
	return BMW{}
}

type MercedesFactory struct{}

func (factory MercedesFactory) CreateCar() Car {
	return Mercedes{}
}

func getCarFactory(model string) CarFactory {
	switch model {
	case "BMW":
		return BMWFactory{}
	case "Mercedes":
		return MercedesFactory{}
	default:
		panic("Unknown car model")
	}
}

func main() {
	// carFactor := getCarFactory("BMW")
	carFactor := getCarFactory("Mercedes")
	car := carFactor.CreateCar()
	car.Drive()
}
