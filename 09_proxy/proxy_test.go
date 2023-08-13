package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProxy(t *testing.T) {
	proxy := Proxy{object: &Object{}}
	require.Equal(t, "before:real:after", proxy.Do())
}
