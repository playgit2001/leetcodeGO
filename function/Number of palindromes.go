package function

import (
	"fmt"
	"strconv"
)

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数 是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 方法1 使用枚举的方式，从前后遍历，判断是否符合.
func IsPalindrome1(x int) bool {
	var strs string = strconv.Itoa(x)
	for i := 0; i < len(strs)/2; i++ {
		if strs[i] != strs[len(strs)-1-i] {
			fmt.Printf("%c  :  %c\n", strs[i], strs[len(strs)-1-i])
			return false
		}
	}
	return true
}
