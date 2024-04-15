package binarysearch

// BinarySearch performs a binary search on a sorted array of integers.
//
// nums is the sorted array of integers, and target is the integer to search for.
// It returns the index of the integer if found, otherwise -1.
func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	low := 0
	high := len(nums)
	for low < high {
		mid := (low + high) / 2

		if target == nums[mid] {
			return mid
		}

		if target < mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
