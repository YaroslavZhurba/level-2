package main

import (
	"fmt"
)

type Ligth struct {
	isOn bool
}

func (l *Ligth) TurnOn() {
	if l.isOn {
		fmt.Println("The light is already switched on")
	} else {
		l.isOn = true
		fmt.Println("Turning on")
	}
}

func (l *Ligth) TurnOff() {
	
	if !l.isOn {
		fmt.Println("The light is already switched off")
	} else {
		l.isOn = false
		fmt.Println("Turning off")
	}
}

type Command interface {
	Execute()
}

type TurnOnLightCommand struct {
	ligth *Ligth
}

func (command TurnOnLightCommand) Execute() {
	command.ligth.TurnOn()
}

type TurnOffLightCommand struct {
	ligth *Ligth
}

func (command TurnOffLightCommand) Execute() {
	command.ligth.TurnOff()
}

type Switch struct {
	flipUpCommand Command
	flipDownCommand Command
}

func (s Switch) FlipUp() {
	s.flipUpCommand.Execute()
}

func (s Switch) FlipDown() {
	s.flipDownCommand.Execute()
}


func main() {
	fmt.Println("Hello")
	ligther := &Ligth{}
	switcher := Switch{
		flipUpCommand: TurnOnLightCommand{ligth: ligther},
		flipDownCommand: TurnOffLightCommand{ligth: ligther},
	}
	switcher.FlipUp()
	switcher.FlipUp()
	switcher.FlipDown()
	switcher.FlipDown()
	switcher.FlipUp()
	switcher.FlipDown()
}