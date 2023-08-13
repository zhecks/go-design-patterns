package proxy

import "fmt"

type ProductProxy struct {
	Order
	IsOwner bool
}

func (pp *ProductProxy) SetID(ID string) error {
	if !pp.IsOwner {
		return fmt.Errorf("no permission")
	}
	pp.Order.SetID(ID)
	return nil
}

func (pp *ProductProxy) SetPrice(price int) error {
	if !pp.IsOwner {
		return fmt.Errorf("no permission")
	}
	pp.Order.SetPrice(price)
	return nil
}

type Order struct {
	ID    string
	Price int
}

func (order *Order) SetID(ID string) {
	order.ID = ID
}

func (order *Order) SetPrice(price int) {
	order.Price = price
}
