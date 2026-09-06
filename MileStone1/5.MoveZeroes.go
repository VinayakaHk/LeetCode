package main

func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		for left < right && nums[left] != 0 {
			left++
		}
		if nums[right] != 0 {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
		}
	}
}
