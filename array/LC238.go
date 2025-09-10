package array

// ProductExceptSelf LC238. 除自己以外的乘积
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
func ProductExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	product := 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			left[i] = 1
			continue
		}
		product = product * nums[i-1]
		left[i] = product
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			product = 1
			right[i] = 1
			continue
		}
		product = product * nums[i+1]
		right[i] = product
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = left[i] * right[i]
	}
	return nums
}
