package iterater

type Aggregate interface {
	CreateIterator()
}

type NumberAggregate struct {
	nums []int
}

func NewNumberAggregate(nums []int) *NumberAggregate {
	if len(nums) == 0 {
		return nil
	}
	return &NumberAggregate{
		nums: nums,
	}
}

func (na *NumberAggregate) CreateIterator() *NumberIterator {
	return &NumberIterator{
		nums: na.nums,
		next: 0,
	}
}

type Iterator interface {
	First()
	Next()
	HasNext() bool
	CurrentItem() interface{}
}

type NumberIterator struct {
	nums []int
	next int
}

func (it *NumberIterator) First() {
	it.next = 0
}

func (it *NumberIterator) Next() {
	if it.next < len(it.nums) {
		it.next++
	}
}

func (it *NumberIterator) HasNext() bool {
	if it.next == len(it.nums) {
		return true
	}
	return false
}

func (it *NumberIterator) CurrentItem() interface{} {
	if !it.HasNext() {
		return it.nums[it.next]
	}
	return nil
}
