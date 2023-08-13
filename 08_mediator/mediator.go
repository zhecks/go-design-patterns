package mediator

import (
	"fmt"
	"strings"
)

var mediator *Mediator

type Mediator struct {
	CD        *CDDriver
	CPU       *CPU
	VideoCard *VideoCard
	SoundCard *SoundCard
}

func GetInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{
			CD:        &CDDriver{},
			CPU:       &CPU{},
			VideoCard: &VideoCard{},
			SoundCard: &SoundCard{},
		}
	}
	return mediator
}

func (mediator *Mediator) Do(typ interface{}) {
	switch instance := typ.(type) {
	case *CDDriver:
		mediator.CPU.Process(instance.Data)
	case *CPU:
		mediator.VideoCard.Display(instance.Video)
		mediator.SoundCard.Play(instance.Sound)
	}
}

type CDDriver struct {
	Data string
}

func (cd *CDDriver) ReadData() {
	cd.Data = "music,image"
	GetInstance().Do(cd)
}

type CPU struct {
	Video string
	Sound string
}

func (cpu *CPU) Process(data string) {
	split := strings.Split(data, ",")
	if len(split) != 2 {
		panic(fmt.Errorf("error data to process"))
	}
	cpu.Sound = split[0]
	cpu.Video = split[1]
	GetInstance().Do(cpu)
}

type VideoCard struct {
	data string
}

func (vc *VideoCard) Display(data string) {
	vc.data = data
}

type SoundCard struct {
	data string
}

func (sc *SoundCard) Play(data string) {
	sc.data = data
}
