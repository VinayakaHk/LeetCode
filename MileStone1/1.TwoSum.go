package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			fmt.Println("1. ", []int{index, i})
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}
