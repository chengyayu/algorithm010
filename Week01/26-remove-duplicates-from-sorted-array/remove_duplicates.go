package test26

func removeDuplicates(nums []int) int {
	if l := len(nums); l <= 1 {
		return l
	}

	p, q := 0, 1
	for q < len(nums) {
		if nums[p] != nums[q] {
			nums[p+1] = nums[q]
			p++
		}
		q++
	}

	return p + 1
}
