package prototype

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestManger(t *testing.T) {
	manger := NewManager()
	manger.Set("order", &OrderApi{})
	prototype := manger.Get("order")
	require.Equal(t, reflect.TypeOf(&OrderApi{}), reflect.TypeOf(prototype))
}
