package heap

type DynamicMedian struct {
	heapMax *Heap
	heapMin *Heap
	size    int
}

func (dm *DynamicMedian) reBalance() {
	if dm.heapMax.Len() > dm.heapMin.Len() {
		dm.heapMin.Insert(dm.heapMax.Remove())
	} else if dm.heapMin.Len() > dm.heapMax.Len() {
		dm.heapMax.Insert(dm.heapMin.Remove())
	}
}

func (dm *DynamicMedian) Len() int {
	return dm.size
}

func (dm *DynamicMedian) Median() int {
	if dm.Len()%2 == 0 {
		return (dm.heapMin.Peek() + dm.heapMax.Peek()) / 2
	} else if dm.heapMax.Len() > dm.heapMin.Len() {
		return dm.heapMax.Peek()
	} else {
		return dm.heapMin.Peek()
	}
}

func (dm *DynamicMedian) Insert(num int) {
	if dm.Len() == 0 {
		dm.heapMin.Insert(num)
		dm.size++
		return
	}

	if num >= dm.Median() {
		dm.heapMin.Insert(num)
	} else {
		dm.heapMax.Insert(num)
	}

	if dm.Len()%2 == 0 {
		dm.reBalance()
	}

	dm.size++
}

// 调试用，非必须
func (dm *DynamicMedian) Drain() *[]int {
	h := make([]int, dm.Len())
	n := dm.heapMax.Len()
	m := dm.heapMin.Len()
	for i := 0; i < n; i++ {
		h[n-i-1] = dm.heapMax.Remove()
	}
	for j := 0; j < m; j++ {
		h[n+j] = dm.heapMin.Remove()
	}
	return &h
}

func construct(data []int, lessFunc func(*[]int, int, int) bool) *Heap {
	h := &Heap{
		data:     data,
		lessFunc: lessFunc,
	}
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	return h
}

func Init(data *[]int) *DynamicMedian {
	d := &DynamicMedian{
		heapMax: construct([]int{}, func(data *[]int, i, j int) bool {
			return (*data)[i] > (*data)[j]
		}),
		heapMin: construct([]int{}, func(data *[]int, i, j int) bool {
			return (*data)[i] < (*data)[j]
		}),
	}
	for _, v := range *data {
		d.Insert(v)
	}
	return d
}
