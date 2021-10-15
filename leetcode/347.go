package leetcode

import (
	"container/heap"
)

/*
 * 347. 前 K 个高频元素
 * 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
 *
 * 示例 1:
 * 	输入: nums = [1,1,1,2,2,3], k = 2
 * 	输出: [1,2]
 *
 * 示例 2:
 * 	输入: nums = [1], k = 1
 * 	输出: [1]
 *
 */
func TopKFrequent(nums []int, k int) []int {
	return topKFrequent(nums, k)
}

type Element struct {
	Num int
	Cnt int
}

type ElementHeap []*Element

func (eh ElementHeap) Len() int           { return len(eh) }
func (eh ElementHeap) Less(i, j int) bool { return eh[i].Cnt < eh[j].Cnt }
func (eh ElementHeap) Swap(i, j int)      { eh[i], eh[j] = eh[j], eh[i] }
func (eh *ElementHeap) Push(x interface{}) {
	*eh = append(*eh, x.(*Element))
}

func (eh *ElementHeap) Pop() interface{} {
	old := *eh
	n := len(old)
	x := old[n-1]
	*eh = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]*Element)
	for _, v := range nums {
		if nil == mp[v] {
			mp[v] = &Element{Num: v, Cnt: 1}
		} else {
			mp[v].Cnt++
		}
	}

	h := &ElementHeap{}
	heap.Init(h)
	for _, v := range mp {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ret := []int{}
	for idx := 0; idx < k && idx < len(mp); idx++ {
		ret = append(ret, heap.Pop(h).(*Element).Num)
	}
	return ret
}
