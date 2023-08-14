package iterater

type PageAggregate struct {
	nums []int
}

func NewPageAggregate(nums []int) *PageAggregate {
	if len(nums) == 0 {
		return nil
	}
	return &PageAggregate{
		nums: nums,
	}
}

func (na *PageAggregate) CreateIterator() *PageIterator {
	return &PageIterator{
		nums: na.nums,
		next: 0,
	}
}

type PageIterator struct {
	nums []int
	next int
}

func (it *PageIterator) First() {
	it.next = 0
}

func (it *PageIterator) HasNext() bool {
	if len(it.nums) != 0 && it.next != len(it.nums) {
		return true
	}
	return false
}

func (it *PageIterator) GetPage(pageNum, pageSize int) []int {
	start, end := (pageNum-1)*pageSize, pageNum*pageSize
	if start < 0 {
		start = 0
	}
	if end > len(it.nums) {
		end = len(it.nums)
	}
	return it.nums[start:end]
}
