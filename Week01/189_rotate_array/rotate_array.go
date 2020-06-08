package _89_rotate_array

//https://leetcode-cn.com/problems/rotate-array/

// 1.暴力法
func rotate1(nums []int, k int) {
	for i := 0; i < k; i++ {
		for j, l := 0, len(nums); j < l; j++ {
			nums[j], nums[l-1] = nums[l-1], nums[j]
		}
	}
}

// 2.环状替换
func rotate2(nums []int, k int) {
	count := 0
	k = k % len(nums)

	for i := 0; count < len(nums); i++ {
		cur := i
		temp := nums[i]

		for {
			next := (cur + k) % len(nums)
			temp, nums[next] = nums[next], temp
			cur = next
			count++
			if i == cur {
				break
			}
		}
	}
}
