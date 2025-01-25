package function

// 暴力枚举，fn1 通过遍历数组的方式，前后相加判断是否等于目标值
func TwoNum1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 第二种方式为利用hashMap的方式，在往map添加元素的过程，判断target-nums[i]是否已经存在于Map中，
// 如果有则给出解，没有则将该元素添加Map，值为该元素下标
func TwoNum2(nums []int, target int) []int {
	length := len(nums)
	numsMap := make(map[int]int)
	for i := 0; i < length; i++ {
		if _, ok := numsMap[target-nums[i]]; ok {
			return []int{numsMap[target-nums[i]], i}
		}
		numsMap[nums[i]] = i
	}
	return nil
}
