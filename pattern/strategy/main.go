package main

import (
	"fmt"
)

type Activity interface {
	DoIt()
}

type Eat struct {}

func (a Eat) DoIt() {
	fmt.Println("Eating...")
}

type Sleep struct {}

func (a Sleep) DoIt() {
	fmt.Println("Sleeping...")
}

type Code struct {}

func (a Code) DoIt() {
	fmt.Println("Coding...")
}

type Programmer struct {
	activity Activity
}

func (p *Programmer) SetActivity(a Activity) {
	p.activity = a
}

func (p Programmer) ExecuteActivity() {
	p.activity.DoIt()
}

func main() {
	programmer := Programmer{}
	programmer.SetActivity(Eat{})
	programmer.ExecuteActivity()
	programmer.SetActivity(Sleep{})
	programmer.ExecuteActivity()
	programmer.SetActivity(Code{})
	programmer.ExecuteActivity()

}