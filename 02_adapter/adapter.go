package adapter

import "strconv"

type Target interface {
	request()
}

type Adaptee struct {
}

func (adaptee Adaptee) specificRequest() int {
	return 1
}

type Adapter struct {
	adaptee Adaptee
}

func NewAdapter(adaptee Adaptee) Adapter {
	return Adapter{
		adaptee,
	}
}

func (adapter Adapter) request() string {
	number := adapter.adaptee.specificRequest()
	return strconv.Itoa(number)
}
