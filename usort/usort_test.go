package usort

import (
	"sort"
	"testing"
)

type Integer []int

func (integer Integer) Len() int {
	return len(integer)
}

func (integer Integer) Less(i, j int) bool {
	return integer[i] < integer[j]
}

func (integer Integer) Swap(i, j int) {
	integer[i], integer[j] = integer[j], integer[i]
}

var integer Integer = []int{39, 99, 93, 69, 11, 26, 89, 62, 46, 72, 2, 14, 77, 86, 35, 64, 55, 30, 66, 18, 75, 28, 68, 67, 81, 21, 20, 58, 10, 24, 51, 1, 27, 48, 85, 97, 25, 92, 6, 23, 34, 36, 15, 80, 9, 3, 41, 22, 90, 61, 70, 73, 16, 17, 71, 57, 31, 94, 47, 13, 65, 12, 78, 56, 82, 88, 40, 45, 76, 8, 38, 91, 4, 100, 83, 49, 19, 32, 53, 84, 63, 42, 5, 87, 43, 29, 7, 33, 95, 74, 60, 79, 98, 59, 54, 50, 44, 96, 37, 52}

func TestBubbleSort(t *testing.T) {
	BubbleSort(integer)
	if !sort.IsSorted(integer) {
		t.Fatal("BubbleSort")
	}
}

func TestSelectionSort(t *testing.T) {
	SelectionSort(integer)
	if !sort.IsSorted(integer) {
		t.Fatal("SelectionSort")
	}
}
