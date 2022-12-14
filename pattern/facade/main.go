package main

import(
	"fmt"
)

type Wheel struct {
	size int
	brand string
}

func (w Wheel) Size() int {
	return w.size
}

func (w Wheel) Brand() string {
	return w.brand}

type Motor struct {
	numberCylinder int
	power float32
}

func (w Motor) StartMotor() {
	fmt.Println("Motor started")
}


type FuelTank struct {
	capacity float32
	size float32
}

func (f FuelTank) SupplyFuel() {
	fmt.Println("Fuel is going to Motor")
}

type Car struct {
	Wheel
	Motor
	FuelTank
}

func (c Car) Start() {
	c.SupplyFuel()
	c.StartMotor()
}

func main() {
	car := &Car{
		Wheel: Wheel{
			size: 16,
			brand: "YOKOHAMA",
		},
		Motor: Motor {
			numberCylinder : 6,
			power : 145,
		},
		FuelTank: FuelTank{
			capacity:  45,
			size : 37.5,
		},
	}
	car.Start()
}