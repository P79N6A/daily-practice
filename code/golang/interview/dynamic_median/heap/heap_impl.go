package heap

type Heap struct {
	data     []int
	lessFunc func(*[]int, int, int) bool
}

func (h *Heap) push(num int) {
	h.data = append(h.data, num)
}
func (h *Heap) pop() int {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}
func (h *Heap) less(i, j int) bool {
	return h.lessFunc(&h.data, i, j)
}
func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *Heap) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}
func (h *Heap) down(i, n int) {
	for {
		j1 := i*2 + 1 // left child
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Peek() int {
	n := 0
	if h.Len() > 0 {
		n = h.data[0]
	}
	return n
}

func (h *Heap) Insert(num int) {
	h.push(num)
	h.up(h.Len() - 1)
}

func (h *Heap) Remove() int {
	n := h.Len() - 1
	h.swap(0, n)
	h.down(0, n)
	return (*h).pop()
}
