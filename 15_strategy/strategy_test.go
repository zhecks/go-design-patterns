package strategy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrategy(t *testing.T) {
	normalCustomerStrategy := NormalCustomerStrategy{}
	oldCustomerStrategy := OldCustomerStrategy{}
	largeCustomerStrategy := LargeCustomerStrategy{}

	price := NewPrice(&normalCustomerStrategy)
	require.Equal(t, float64(1000), price.quoto(1000))
	price = NewPrice(&oldCustomerStrategy)
	require.Equal(t, float64(950), price.quoto(1000))
	price = NewPrice(&largeCustomerStrategy)
	require.Equal(t, float64(900), price.quoto(1000))
}
