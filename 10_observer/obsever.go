package observer

import "fmt"

type Subject struct {
	Observers []Observer
	Context   string
}

func (subject *Subject) Attach(observer Observer) {
	subject.Observers = append(subject.Observers, observer)
}

func (subject *Subject) UpdateContext(context string) {
	subject.Context = context
	subject.notifyAll()
}

func (subject *Subject) notifyAll() {
	for _, observer := range subject.Observers {
		observer.Update(subject)
	}
}

type Observer struct {
	Name string
}

func (o *Observer) Update(subject *Subject) {
	fmt.Printf("%s receive %s\n", o.Name, subject.Context)
}
