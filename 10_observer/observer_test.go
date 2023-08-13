package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := Subject{}
	subject.Attach(Observer{"observer1"})
	subject.Attach(Observer{"observer2"})
	subject.UpdateContext("context")
}
