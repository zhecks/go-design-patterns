package interpreter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInterpreter(t *testing.T) {
	parser := Parser{}
	parser.Parse("1 + 2 + 3 - 4 + 5 - 6")
	require.Equal(t, 1, parser.Result())
}
