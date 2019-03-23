package ints

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(v interface{}) { *h = append(*h, v.(int)) }

func (h *Heap) Pop() (v interface{}) {
	v = (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return
}
