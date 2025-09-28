package daily

// MaxFrequencyElements LC3005 最大频率元素计数
// 给你一个由 正整数 组成的数组 nums 。
// 返回数组 nums 中所有具有 最大 频率的元素的 总频率 。
// 元素的 频率 是指该元素在数组中出现的次数。
func MaxFrequencyElements(nums []int) int {
	// 1. 法一：直接用map来存数和频率
	/*	freqMap := map[int]int{}

		maxFreq := 1
		for _, num := range nums {
			freq, ok := freqMap[num]
			if ok {
				freqMap[num] = freq + 1
				if freq >= maxFreq {
					maxFreq = freq + 1
				}
			} else {
				freqMap[num] = 1
			}
		}

		count := 0
		for _, v := range freqMap {
			if v == maxFreq {
				count++
			}
		}
		return count * maxFreq*/

	// 2. 法二：直接用数组，因为题目提示nums[i]在0~100之间
	freq := make([]int, 101, 101)
	for _, num := range nums {
		freq[num]++
	}

	res := 0
	maxNum := 0

	for _, num := range freq {
		if num > maxNum {
			res = num
			maxNum = num
		} else if num == maxNum {
			res = res + maxNum
		}
	}
	return res
}
