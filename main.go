package main

import (
	"fmt"
	"go-algorithm/leetcode"
	"go-algorithm/pattern/strategy"
	"go-algorithm/pattern/template"
	"go-algorithm/usort"
	"math/rand"
	"sort"
	"time"
	"unicode/utf8"
)

var arr = []int{1, 2, 3, 4, 11, 6, 11, 12, 13, 14, 15, 7, 8, 9, 10}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	for idx := 0; idx < 10; idx++ {
		fmt.Println(leetcode.BinarySearch([]int{1, 3, 5, 6}, idx))
	}
	return

	fmt.Println(leetcode.FindSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	return

	fmt.Println(leetcode.SingleNumber([]int{1, 3, 1, -1, 3}))
	return

	fmt.Println(arr)
	leetcode.MergeSort(arr)
	leetcode.QuickSort(arr)
	fmt.Println(arr)
	return

	//matrix_0 := [][]byte{[]byte{'1', '0', '1', '0', '0'}, []byte{'1', '0', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1'}, []byte{'1', '0', '0', '1', '0'}}
	matrix_0 := [][]byte{[]byte{'0', '1'}, []byte{'0', '1'}}
	fmt.Println(leetcode.MaximalRectangle(matrix_0))
	return
	fmt.Println(leetcode.MaxSlidingWindow([]int{1, 3, -1, -3, -4, 3, 6, 7}, 3))
	return

	fmt.Println(leetcode.LargestRectangleArea([]int{4, 2}))
	return

	narr := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	tmp := [][]int{}

	tmp = append(tmp, narr[0])
	fmt.Println(tmp)
	fmt.Println(narr)

	narr[0][1] = 9999
	fmt.Println(tmp)
	fmt.Println(narr)
	return

	fmt.Println(leetcode.Convert("PAYPALISHIRING", 4))
	return

	bg := make([]byte, 20)
	bg1 := bg[0:0:10]
	bg2 := bg[10:10]
	bg1 = append(bg1, 'a')
	bg2 = append(bg2, 'b')
	bg1 = append(bg1, 'c')
	bg2 = append(bg2, 'd')
	fmt.Println(len(bg), string(bg))
	fmt.Println(len(bg1), string(bg1))
	fmt.Println(len(bg2), string(bg2))
	return

	//outerArr := []int{1, 2, 3}
	//for k, v := range outerArr {
	//	outerArr[0], outerArr[k] = outerArr[k], outerArr[0]
	//	fmt.Println(k, v, outerArr)
	//}
	//fmt.Println(outerArr)
	//return

	fmt.Println(leetcode.PermuteUnique([]int{1, 1, 2}))
	return

	fmt.Println(leetcode.Permute([]int{1, 2}))
	return

	fmt.Println(leetcode.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	return

	fmt.Println(leetcode.TwoSum([]int{2, 7, 11, 15}, 9))
	return

	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)
	return
	//fmt.Println(leetcode.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(leetcode.MaxArea([]int{}))
	return

	fmt.Println(leetcode.LongestPalindrome("babad"))
	return

	mp := map[int]int{99: 3}
	v := mp[3]
	fmt.Println(v)

	return

	str := "abcdefghijklmnopqrstuvwxyz"
	for k, v := range str {
		fmt.Printf("%d\t%c\n", k, v)
	}
	return

	//fmt.Println(leetcode.Subsets([]int{1, 2, 3, 8}))
	fmt.Println(leetcode.Subsets([]int{9, 0, 3, 5, 1}))
	return

	combine_result := leetcode.Combine(6, 2)
	fmt.Println(combine_result, len(combine_result))
	return

	colors := []int{2, 0, 2, 1, 1, 0}
	leetcode.SortColors(colors)
	fmt.Println(colors)
	return

	fmt.Println(leetcode.SearchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 1))
	return
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	leetcode.SetZeroes(matrix)
	fmt.Println(matrix)
	return

	fmt.Println(leetcode.FrequencySort("cccaaa我的名字哈哈哈哈哈"))
	return
	fmt.Println(leetcode.FrequencySort("cccaaa我的名字哈哈哈哈哈"))
	return

	bigArr := []int{}
	for idx := 9999999; idx >= 0; idx-- {
		bigArr = append(bigArr, rand.Intn(99999999))
	}
	//usort.HeapSort(bigArr)
	//usort.MergeSort(bigArr)
	usort.QuickSort(bigArr)
	//usort.ShellSort(bigArr)
	//sort.Ints(bigArr)
	fmt.Println(sort.IntsAreSorted(bigArr))
	fmt.Println(bigArr[9999970:9999999])
	return

	fmt.Println(utf8.RuneError)
	fmt.Println('\uFFFD')
	fmt.Println(utf8.RuneSelf)
	fmt.Println(utf8.MaxRune)
	fmt.Println(utf8.UTFMax)

	leetcode.Arrangement("abcdefg", 3)
	leetcode.Arrangement("中国", 3)
	leetcode.Arrangement("ab中国", 3)
	return

	fmt.Println(leetcode.FindRepeatNumber(arr))
	return

	fmt.Println(-1 << 2)
	fmt.Println(leetcode.IsHappyNumber(9))
	return

	//arr := make([]int, 0)
	for idx := 0; idx <= 10; idx++ {
		arr = append(arr, 500-rand.Intn(1000))
	}
	fmt.Println(leetcode.MaxSubArr(arr))
	fmt.Println(leetcode.MaxSubArr1(arr))
	return

	fishsoup := template.NewFishSoup()
	fishsoup.Cooking()

	//fmt.Println()
	barbecue := template.NewBarbecue()
	barbecue.Cooking()
	return

	rect := strategy.Rect{3.2, 89.1}
	circle := strategy.Circle{8.1}

	shapeList := []strategy.Shape{rect, circle}
	fmt.Println(shapeList)

	for _, shape := range shapeList {
		fmt.Println(shape.Area())
		fmt.Println(shape.Perimeter())
	}
	fmt.Println()

	context := strategy.Context{rect}
	fmt.Println(context.GetArea())
	fmt.Println(context.GetPerimeter())
	return

	fmt.Println(arr)
	//usort.QuickSort_1(arr)
	usort.HeapSort_1(arr)
	fmt.Println(arr)
	return
	usort.MergeSort_1(arr)
	fmt.Println(arr)
	usort.ShellSort_1(arr)
	fmt.Println(arr)
	usort.InsertionSort_1(arr)
	fmt.Println(arr)
	usort.BubbleSort_1(arr)
	fmt.Println(arr)
	usort.SelectionSort_1(arr)
	fmt.Println(arr)

	/*
		usort.HeapSort(arr)
		fmt.Println(arr)
		return
		usort.MergeSort(arr)
		fmt.Println(arr)
		return
		usort.ShellSort(arr)
		fmt.Println(arr)
		usort.QuickSort(arr)
		fmt.Println(arr)
		usort.BubbleSort(arr)
		fmt.Println(arr)
		usort.SelectionSort(arr)
		fmt.Println(arr)
		usort.InsertionSort(arr)
		fmt.Println(arr)
	*/
}
