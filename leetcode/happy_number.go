package leetcode

import (
	"fmt"
)

func IsHappyNumber(number int) bool {

	fmt.Println(number)
	if number == 1 {
		return true
	}

	if number == 4 {
		return false
	}

	squareSum := 0
	for number > 0 {
		num := number % 10
		number /= 10

		squareSum += num * num
	}

	return IsHappyNumber(squareSum)
}
