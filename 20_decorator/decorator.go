package decorator

type Component interface {
	Calc() int
}

type ComponentImpl struct {
}

func (impl *ComponentImpl) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.num * d.Component.Calc()
}

type AddDecorator struct {
	Component
	num int
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.num + d.Component.Calc()
}
