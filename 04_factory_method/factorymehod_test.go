package factorymethod

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExportTxtOperate(t *testing.T) {
	factory := ExportTxtFactory{}
	operate := factory.create()
	require.Equal(t, "father", operate.father())
	require.Equal(t, "txt export operate", operate.export())
}

func TestExportTomlOperate(t *testing.T) {
	factory := ExportTomlFactory{}
	operate := factory.create()
	require.Equal(t, "father", operate.father())
	require.Equal(t, "toml export operate", operate.export())
}
