package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	fmt.Println("3. ", nums[len(nums)/2])
	return nums[len(nums)/2]
}
