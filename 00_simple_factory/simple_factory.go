package simplefactory

import "fmt"

type Connector interface {
	conn() string
}

var _ Connector = (*mysql)(nil)

type mysql struct {
}

func (i mysql) conn() string {
	return fmt.Sprintf("mysql connecting...")
}

var _ Connector = (*oracle)(nil)

type oracle struct {
}

func (i oracle) conn() string {
	return fmt.Sprintf("oracle connecting...")
}

func CreateDb(typ int) Connector {
	switch typ {
	case 1:
		return mysql{}
	case 2:
		return oracle{}
	default:
		return nil
	}
}
