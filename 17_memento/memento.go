package memento

import "fmt"

type FlowAMock struct {
	FlowName   string
	tempResult int
	tempState  string
}

func NewFlowAMock(flowName string) *FlowAMock {
	return &FlowAMock{
		FlowName: flowName,
	}
}

func (flow *FlowAMock) runPhaseOne() {
	flow.tempResult = 3
	flow.tempState = "PhaseOne"
}

func (flow *FlowAMock) schema1() {
	flow.tempState += ", Schema1"
	fmt.Println(flow.tempState+": now run ", flow.tempResult)
	flow.tempResult += 11
}

func (flow *FlowAMock) schema2() {
	flow.tempState += ", Schema2"
	fmt.Println(flow.tempState+": now run ", flow.tempResult)
	flow.tempResult += 22
}

func (flow *FlowAMock) CreateMemento() FlowAMockMemento {
	return &flowAMockMementoImpl{
		tempResult: flow.tempResult,
		tempState:  flow.tempState,
	}
}

func (flow *FlowAMock) SetMemento(memento FlowAMockMemento) {
	impl := memento.(*flowAMockMementoImpl)
	flow.tempResult = impl.tempResult
	flow.tempState = impl.tempState
}

type FlowAMockMemento interface {
}

type FlowAMementoCareTaker struct {
	memento FlowAMockMemento
}

func (ct *FlowAMementoCareTaker) SaveMemento(memento FlowAMockMemento) {
	ct.memento = memento
}

func (ct *FlowAMementoCareTaker) RetriveMemento() FlowAMockMemento {
	return ct.memento
}

type flowAMockMementoImpl struct {
	tempResult int
	tempState  string
}

func NewFlowAMockMementoImpl(tempResult int, tempState string) *flowAMockMementoImpl {
	return &flowAMockMementoImpl{
		tempResult: tempResult,
		tempState:  tempState,
	}
}

func (impl *flowAMockMementoImpl) GetTempResult() int {
	return impl.tempResult
}

func (impl *flowAMockMementoImpl) GetTempState() string {
	return impl.tempState
}
