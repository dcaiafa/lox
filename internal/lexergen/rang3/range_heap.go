package rang3

import (
	"container/heap"
)

type rangeHeapInternal []Range

func (h rangeHeapInternal) Len() int { return len(h) }

func (h rangeHeapInternal) Less(i, j int) bool {
	return Compare(h[i], h[j]) < 0
}

func (h rangeHeapInternal) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *rangeHeapInternal) Push(x any) {
	item := x.(Range)
	*h = append(*h, item)
}

func (h *rangeHeapInternal) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type rangeHeap struct {
	heap rangeHeapInternal
	set  map[Range]bool
}

func newRangeHeap(ranges []Range) *rangeHeap {
	rh := &rangeHeap{
		heap: make(rangeHeapInternal, 0, len(ranges)),
		set:  make(map[Range]bool, len(ranges)),
	}

	for _, r := range ranges {
		rh.Push(r)
	}

	return rh
}

func (rh *rangeHeap) Push(r Range) {
	if rh.set[r] {
		return
	}
	rh.set[r] = true
	heap.Push(&rh.heap, r)
}

func (rh *rangeHeap) Pop() Range {
	r := rh.heap[0]
	delete(rh.set, r)
	heap.Pop(&rh.heap)
	return r
}

func (rh *rangeHeap) Len() int {
	return len(rh.heap)
}

func (rh *rangeHeap) Peek() Range {
	return rh.heap[0]
}
