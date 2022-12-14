package main

import (
	"fmt"
)

type Food interface {
	visit(FoodVisitor)
}

type Mustard struct {
	taste int
}

func (mustard Mustard) visit(foodVisitor FoodVisitor) {
	foodVisitor.onMustard(mustard)
}

type Juise struct {
	taste int
}

func (juise Juise) visit(foodVisitor FoodVisitor) {
	foodVisitor.onJuise(juise)
}

type Cake struct {
	taste int
}

func (cake Cake) visit(foodVisitor FoodVisitor) {
	foodVisitor.onCake(cake)
}

type FoodVisitor struct{}

func (fv FoodVisitor) printTaste(taste int) {
	if taste == 0 {
		fmt.Print("It's awful!!!")
	} else if taste == 1 {
		fmt.Print("It's pretty...")
	} else if taste == 2 {
		fmt.Print("It's GREAT :)")
	}
}

func (fv FoodVisitor) onMustard(mustard Mustard) {
	fv.printTaste(mustard.taste)
	fmt.Println(" mustard")
}

func (fv FoodVisitor) onJuise(juise Juise) {
	fv.printTaste(juise.taste)
	fmt.Println(" juise")
}

func (fv FoodVisitor) onCake(cake Cake) {
	fv.printTaste(cake.taste)
	fmt.Println(" cake")
}

func main() {
	meal := []Food{Mustard{taste: 0}, Juise{taste: 1}, Cake{taste: 2}}
	foodVisitor := FoodVisitor{}

	for _, food := range meal {
		food.visit(foodVisitor)
	}
}
