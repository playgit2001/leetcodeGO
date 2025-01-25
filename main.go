package main

import (
	"leetcode/function"
)

func main() {
	var nums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var result []int
	result = function.TwoNum2(nums, 11)
	for _, v := range result {
		println(v + 1)
	}
}
