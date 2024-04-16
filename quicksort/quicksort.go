package quicksort

func QuickSortRecursive(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}

	pivot := nums[len(nums)-1]

	left, right := partition(nums, pivot)
	return append(append(QuickSortRecursive(left), pivot), QuickSortRecursive(right)...)
}

func partition(nums []int, pivot int) ([]int, []int) {
	l := make([]int, 0)
	r := make([]int, 0)
	for _, n := range nums {
		if n < pivot {
			l = append(l, n)
		}
		if n > pivot {
			r = append(r, n)
		}
	}

	return l, r
}
