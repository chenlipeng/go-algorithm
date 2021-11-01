package leetcode

import (
	"container/heap"
)

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(e interface{}) {
	*h = append(*h, e.(int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func FindKthLargest(nums []int, k int) int {
	return findKthLargest(nums, k)
}

func findKthLargest(nums []int, k int) int {
	hp := &IHeap{}
	heap.Init(hp)
	for _, v := range nums {
		heap.Push(hp, v)
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}

	return heap.Pop(hp).(int)
}
