package array

// Rotate 轮换数组
// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func Rotate(nums []int, k int) []int {
	k = k % len(nums)
	// 1. dig k nums
	lastK := make([]int, k)
	copy(lastK, nums[len(nums)-1-k+1:])
	// 2.
	for i := len(nums) - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	copy(nums, lastK)
	return nums
}
