package mergesort

func MergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	return MergeArrays(MergeSort(nums[:mid]), MergeSort(nums[mid:]))
}

func MergeArrays(a []int, b []int) []int {
	merged := make([]int, 0)
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		merged = append(merged, a[i])
	}
	for ; j < len(b); j++ {
		merged = append(merged, b[j])
	}

	return merged
}
