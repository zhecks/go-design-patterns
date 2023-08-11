package singleton

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	require.Equal(t, "singleton foo", instance1.foo())
	instance2 := GetInstance()
	require.Equal(t, instance1, instance2)
}
