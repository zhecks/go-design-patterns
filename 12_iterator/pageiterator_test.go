package iterater

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPageIterator(t *testing.T) {
	aggregate := NewPageAggregate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	iterator := aggregate.CreateIterator()
	iterator.First()
	page := iterator.GetPage(2, 2)
	require.Equal(t, []int{3, 4}, page)
}
