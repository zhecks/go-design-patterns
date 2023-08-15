package decorator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecorator(t *testing.T) {
	impl := &ComponentImpl{}
	require.Equal(t, 4, WrapMulDecorator(WrapAddDecorator(impl, 2), 2).Calc())
}
