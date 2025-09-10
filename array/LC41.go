package array

// FirstMissingPositive LC41. 确实的第一个正数
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
func FirstMissingPositive(nums []int) int {
	// 思路1
	/*	sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})

		p := -1
		for i := 0; i < len(nums); i++ {
			if nums[i] > 0 {
				p = i
				break
			}
		}
		if p < 0 {
			return 1
		}

		for i := p; i < len(nums); i++ {
			if i == 0 {
				if nums[i] > 1 {
					return 1
				}
				continue
			}
			if nums[i]-1 > nums[i-1] {
				switch {
				case nums[i-1] >= 0:
					return nums[i-1] + 1
				case nums[i]-1 >= 1:
					return 1
				default:
					continue
				}

			}

		}
		return nums[len(nums)-1] + 1*/

	// 思路2
	/*	sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		minNum := 1
		for i := 0; i < len(nums); i++ {
			if nums[i] == minNum {
				minNum++
			}
		}
		return minNum*/

	// 思路3 时间复杂度O(n)
	records := make([]int, len(nums)+1)
	for _, num := range nums {
		if num > len(nums) || num <= 0 {
			continue

		}
		records[num] = 1
	}
	for i := 1; i < len(records); i++ {
		if records[i] == 1 {
			continue
		}
		return i
	}
	return len(nums) + 1

}
