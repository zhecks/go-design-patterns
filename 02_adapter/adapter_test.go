package adapter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequest(t *testing.T) {
	adaptee := Adaptee{}
	adapter := Adapter{adaptee: adaptee}
	require.Equal(t, "1", adapter.request())
}
