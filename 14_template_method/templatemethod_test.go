package templatemethod

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTemplate(t *testing.T) {
	lm := LoginMode{
		LoginID:  "1",
		Password: "123",
	}
	normalLogin := NewTemplate(&NormalLogin{})
	require.True(t, normalLogin.Login(lm))
	wokerLogin := NewTemplate(&WokerLogin{})
	require.True(t, wokerLogin.Login(lm))
}
