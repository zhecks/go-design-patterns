package prototype

type Cloneable interface {
	clone() Cloneable
}

type Manger struct {
	prototypes map[string]Cloneable
}

func NewManager() *Manger {
	return &Manger{
		prototypes: map[string]Cloneable{},
	}
}

func (m *Manger) Get(name string) Cloneable {
	return m.prototypes[name].clone()
}

func (m *Manger) Set(name string, cloneable Cloneable) {
	m.prototypes[name] = cloneable
}

var _ Cloneable = (*OrderApi)(nil)

type OrderApi struct {
}

func (o *OrderApi) clone() Cloneable {
	return &OrderApi{}
}
