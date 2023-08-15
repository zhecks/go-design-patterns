package memento

import "testing"

func TestMemento(t *testing.T) {
	mock := NewFlowAMock("TestFlow")
	mock.runPhaseOne()
	memento := mock.CreateMemento()
	taker := FlowAMementoCareTaker{}
	taker.SaveMemento(memento)

	mock.schema1()
	mock.SetMemento(taker.RetriveMemento())
	mock.schema2()
}
