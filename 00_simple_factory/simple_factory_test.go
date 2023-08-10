package simplefactory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMysql(t *testing.T) {
	db := CreateDb(1)
	require.Equal(t, "mysql connecting...", db.conn())
}

func TestOracle(t *testing.T) {
	db := CreateDb(2)
	require.Equal(t, "oracle connecting...", db.conn())
}
