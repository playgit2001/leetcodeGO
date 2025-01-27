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

//第二种，反转数字法，不通过转字符串的方式，采用反转一半数字的方式。
//比如说，1221，我们截取21，把他构造成12，如果首12=构造12，那么这个是回文数
//首先是排除边界情况，个位数是0以及负数不可能是回文数。
//我们先把数字%10，余数不断乘10，累加上去，直到取了一半数。
//判断是否到达一半的方式，则是构造的数是否等于被不断整除后的原始数

func IsPalindrome2(x int) bool {
	if x < 0 || x%10 == 0 {
		return false
	}
	var res int = 0
	for !(res == x || res > x) {
		res = res*10 + x%10
		x = x / 10
		fmt.Println("%c ,%c", res, x)
	}
	if res == x {
		return true
	}

	if res > x {
		if res/10 == x {
			return true
		}
	}
	return false
}
