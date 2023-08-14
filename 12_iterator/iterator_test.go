package iterater

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	aggregate := NewNumberAggregate([]int{1, 2, 3})
	iterator := aggregate.CreateIterator()
	iterator.First()
	for !iterator.HasNext() {
		fmt.Println(iterator.CurrentItem())
		iterator.Next()
	}
}
