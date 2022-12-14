package main

import (
	"fmt"
	"reflect"
)

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

type Activity interface {
	DoIt()
}

type Eat struct{}

func (a Eat) DoIt() {
	fmt.Println("Eating...")
}

type Sleep struct{}

func (a Sleep) DoIt() {
	fmt.Println("Sleeping...")
}

type Code struct{}

func (a Code) DoIt() {
	fmt.Println("Coding...")
}

type Programmer struct {
	activity Activity
}

func (p *Programmer) ChangeActivity() {
	if IsInstanceOf(p.activity, Eat{}) {
		p.activity = Sleep{}
	} else if IsInstanceOf(p.activity, Sleep{}) {
		p.activity = Code{}
	} else if IsInstanceOf(p.activity, Code{}) {
		p.activity = Eat{}
	}
}

func (p Programmer) ExecuteActivity() {
	p.activity.DoIt()
}

func main() {
	programmer := Programmer{activity: Eat{}}
	programmer.ExecuteActivity()
	programmer.ChangeActivity()
	programmer.ExecuteActivity()
	programmer.ChangeActivity()
	programmer.ExecuteActivity()
}
