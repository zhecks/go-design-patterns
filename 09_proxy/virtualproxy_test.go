package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVirtualProxy(t *testing.T) {
	manager := UserManager{}
	user := manager.GetUser()
	require.Equal(t, "1", user.GetID())
	require.Equal(t, "Alice", user.GetName())
	require.Equal(t, "detail", user.GetDetail())
}
