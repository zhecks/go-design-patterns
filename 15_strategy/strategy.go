package strategy

import "fmt"

type Strategy interface {
	calcPrice(goodsPrice float64) float64
}

type NormalCustomerStrategy struct {
	Strategy
}

func (strategy *NormalCustomerStrategy) calcPrice(goodsPrice float64) float64 {
	fmt.Println("对于新客户或者普通客户，没有折扣")
	return goodsPrice
}

type OldCustomerStrategy struct {
	Strategy
}

func (strategy *OldCustomerStrategy) calcPrice(goodsPrice float64) float64 {
	fmt.Println("对于老客户，统一折扣5%")
	return goodsPrice * (1 - 0.05)
}

type LargeCustomerStrategy struct {
	Strategy
}

func (strategy *LargeCustomerStrategy) calcPrice(goodsPrice float64) float64 {
	fmt.Println("对于大客户，统一折扣10%")
	return goodsPrice * (1 - 0.1)
}

type Price struct {
	strategy Strategy
}

func NewPrice(strategy Strategy) *Price {
	return &Price{strategy}
}

func (price *Price) quoto(goodsPrice float64) float64 {
	return price.strategy.calcPrice(goodsPrice)
}
