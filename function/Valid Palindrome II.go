package function

// 验证回文串2
// 给你一个字符串 s，最多 可以从中删除一个字符。
// 请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。

// 这道题我能理解一部分，但是不是很深刻
// 1.使用双指针法，为什么使用双指针而不是length/2，因为这道题题的中间点，在排除了不同点后，这个中间点是会变化的。
// 2.循环条件为左指针小于右指针，如果遇到不同点，分为2部分，一个是排除左边点，左指针+1，一个是排除右边点，右指针-1，当这两种情况有一个能得到回文串，那么就可以在最多只删除一个点的情况下，获得回文串。
func ValidPalindromeII(strs string) bool {
	length := len(strs)
	var (
		left  int = 0
		right int = length - 1
	)
	for left < right {
		if strs[left] != strs[right] {
			return tmp(strs, left+1, right) || tmp(strs, left, right-1)
		}
		left++
		right--
	}
	return true
}
func tmp(strs string, left, right int) bool {
	for left < right {
		if strs[left] != strs[right] {
			return false
		}
		left++
		right--
	}
	return true
}
