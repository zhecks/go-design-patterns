package abstractfactory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPCFactory(t *testing.T) {
	factory := PCFactory{}
	require.Equal(t, "PC productA", factory.createProductA())
	require.Equal(t, "PC productB", factory.createProductB())
}

func TestPhoneFactory(t *testing.T) {
	factory := PhoneFactory{}
	require.Equal(t, "Phone productA", factory.createProductA())
	require.Equal(t, "Phone productB", factory.createProductB())
}
