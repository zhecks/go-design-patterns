package facade

import "fmt"

type AModuleImpl struct {
}

func (impl AModuleImpl) testA(number int) int {
	fmt.Println("do module A...")
	return number + 1
}

type BModuleImpl struct {
}

func (impl BModuleImpl) testB(number int) int {
	fmt.Println("do module B...")
	return number + 2
}

type CModuleImpl struct {
}

func (impl CModuleImpl) testC(number int) int {
	fmt.Println("do module C...")
	return number + 3
}

type Facade struct {
}

func (facade Facade) test(number int) int {
	a := AModuleImpl{}
	number = a.testA(number)
	b := BModuleImpl{}
	number = b.testB(number)
	c := CModuleImpl{}
	number = c.testC(number)
	return number
}
