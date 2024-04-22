package kadane

import (
	"math"
)

func MaxSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}

	// sub := make([]int, 0)
	maxSum := -math.MaxInt
	currSum := 0
	maxLeft := 0
	maxRight := 0
	left := 0
	for i, n := range nums {
		if currSum < 0 {
			currSum = 0
			left = i
		}

		currSum += n

		if currSum > maxSum {
			maxSum = currSum
			maxLeft, maxRight = left, i
		}
	}

	return nums[maxLeft : maxRight+1]
}
