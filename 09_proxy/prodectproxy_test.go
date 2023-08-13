package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductProxy(t *testing.T) {
	proxy := ProductProxy{Order{}, true}
	err := proxy.SetID("1")
	require.Nil(t, err)
	require.Equal(t, "1", proxy.ID)
	err = proxy.SetPrice(100)
	require.Nil(t, err)
	require.Equal(t, 100, proxy.Price)
}

func TestProductProxy_NoPermission(t *testing.T) {
	proxy := ProductProxy{Order{}, false}
	err := proxy.SetID("1")
	require.Equal(t, "no permission", err.Error())
	err = proxy.SetPrice(100)
	require.Equal(t, "no permission", err.Error())
}
