package facade

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTest(t *testing.T) {
	facade := Facade{}
	require.Equal(t, 6, facade.test(0))
}
