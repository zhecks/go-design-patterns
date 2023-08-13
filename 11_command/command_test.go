package command

import (
	"testing"
)

func TestCommand(t *testing.T) {
	buttonA := &StartCommand{}
	buttonB := &RebootComand{}
	box := NewBox(buttonA, buttonB)
	box.PressButtonA()
	box.PressButtonB()
}
