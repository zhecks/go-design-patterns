package builder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildXml(t *testing.T) {
	director := Director{build: &XmlBuilder{}}
	require.Equal(t, "XmlHeaderXmlBodyXmlFooter", director.construct())
}

func TestBuildTxt(t *testing.T) {
	director := Director{build: &TxtBuilder{}}
	require.Equal(t, "TxtHeaderTxtBodyTxtFooter", director.construct())
}
