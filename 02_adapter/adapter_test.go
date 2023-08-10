package adapter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRequest(t *testing.T) {
	adaptee := Adaptee{}
	adapter := Adapter{adaptee: adaptee}
	require.Equal(t, "1", adapter.request())
}
