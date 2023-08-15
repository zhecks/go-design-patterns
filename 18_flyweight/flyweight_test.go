package flyweight

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlyweight(t *testing.T) {
	instance := GetInstance()
	image1 := instance.Get("image1")
	image2 := instance.Get("image1")
	require.Equal(t, image2.Data(), image1.Data())
	require.Equal(t, image2, image1)
}
