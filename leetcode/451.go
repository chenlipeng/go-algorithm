package leetcode

import (
	"sort"
)

func FrequencySort(s string) string {
	return frequencySort_1(s)
}

//解法二、、、、、、、
func frequencySort_1(s string) string {
	mp := map[rune]int{}
	for _, v := range s {
		mp[v] += 1
	}

	type pair struct {
		ch  rune
		cnt int
	}

	pairs := []pair{}
	for k, v := range mp {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })

	ret := []rune{}
	for _, v := range pairs {
		for idx := 0; idx < v.cnt; idx++ {
			ret = append(ret, v.ch)
		}
	}

	return string(ret)
}

//解法一、、、、、、、

func frequencySort(s string) string {
	mp := map[rune]int{}
	//每个元素的个数
	for _, v := range s {
		mp[v] += 1
	}

	//放入map中
	var list NodeSlice
	for k, v := range mp {
		list = append(list, &Node{
			V: k,
			N: v,
		})
	}
	//排序
	sort.Sort(list)

	var ret []rune
	for _, v := range list {
		for idx := 0; idx < v.N; idx++ {
			ret = append(ret, v.V)
		}
	}
	return string(ret)
}

type Node struct {
	V rune
	N int
}

type NodeSlice []*Node

func (ns NodeSlice) Len() int {
	return len(ns)
}

func (ns NodeSlice) Less(i, j int) bool {
	return ns[i].N >= ns[j].N
}

func (ns NodeSlice) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}
