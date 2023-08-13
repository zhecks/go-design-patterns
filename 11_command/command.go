package command

import "fmt"

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func (cmd *StartCommand) Execute() {
	cmd.mb.Start()
}

type RebootComand struct {
	mb *MotherBoard
}

func (cmd *RebootComand) Execute() {
	cmd.mb.Reboot()
}

type MotherBoard struct {
}

func (mb *MotherBoard) Start() {
	fmt.Println("system starting...")
}

func (mb *MotherBoard) Reboot() {
	fmt.Println("system rebooting...")
}

type Box struct {
	buttonA Command
	buttonB Command
}

func NewBox(buttonA, buttonB Command) *Box {
	return &Box{
		buttonA: buttonA,
		buttonB: buttonB,
	}
}

func (box *Box) PressButtonA() {
	box.buttonA.Execute()
}

func (box *Box) PressButtonB() {
	box.buttonB.Execute()
}
