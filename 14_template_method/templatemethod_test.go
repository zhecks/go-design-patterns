package templatemethod

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTemplate(t *testing.T) {
	lb := LoginMode{
		LoginID:  "1",
		Password: "123",
	}
	dblb1 := LoginMode{
		LoginID:  "1",
		Password: "123",
	}
	dblb2 := LoginMode{
		LoginID:  "1",
		Password: "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	}
	normalLogin := NormalLogin{}
	require.True(t, normalLogin.Match(lb, dblb1))
	wokerLogin := WokerLogin{}
	require.True(t, wokerLogin.Match(lb, dblb2))
}
