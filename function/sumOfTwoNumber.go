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

//测试用例
//func main() {
//	var nums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	var result []int
//	result = function.TwoNum2(nums, 11)
//	for _, v := range result {
//		println(v + 1)
//	}
//}
