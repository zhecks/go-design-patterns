package singleton

import "sync"

type Singleton interface {
	foo() string
}

type singleton struct {
}

func (st *singleton) foo() string {
	return "singleton foo"
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
