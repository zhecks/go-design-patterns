package chainofresponsibility

import (
	"fmt"
	"reflect"
)

type Handler interface {
	HaveRight(fee float64) bool
	HandleFeeRequest(fee float64) bool
}

type RequestChain struct {
	Handler
	successor *RequestChain
}

func (r *RequestChain) SetSuccessor(successor *RequestChain) {
	r.successor = successor
}

func (r *RequestChain) HandleFeeRequest(fee float64) bool {
	if r.Handler.HaveRight(fee) {
		typ := reflect.TypeOf(r.Handler)
		fmt.Printf("handle fee by %s\n", typ.String())
		return r.Handler.HandleFeeRequest(fee)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(fee)
	}
	return false
}

func (r *RequestChain) HaveRight(_ float64) bool {
	return true
}

type GeneralMananger struct {
}

func NewGeneralManagerRequest() *RequestChain {
	return &RequestChain{
		Handler:   &GeneralMananger{},
		successor: nil,
	}
}

func (manager *GeneralMananger) HaveRight(fee float64) bool {
	return fee <= 500
}

func (manager *GeneralMananger) HandleFeeRequest(_ float64) bool {
	return true
}

type ProjectMananger struct {
}

func NewProjectManagerRequest() *RequestChain {
	return &RequestChain{
		Handler:   &ProjectMananger{},
		successor: nil,
	}
}

func (manager *ProjectMananger) HaveRight(fee float64) bool {
	return fee <= 1000
}

func (manager *ProjectMananger) HandleFeeRequest(_ float64) bool {
	return true
}

type DepMananger struct {
}

func NewDepManagerRequest() *RequestChain {
	return &RequestChain{
		Handler:   &DepMananger{},
		successor: nil,
	}
}

func (manager *DepMananger) HaveRight(_ float64) bool {
	return true
}

func (manager *DepMananger) HandleFeeRequest(_ float64) bool {
	return true
}
