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

func main() {
	a := []int{2, 7, 11, 15}
	target := 9
	twoSum(a, target)
	bestTimeToBuyAndSellStock(a)
	majorityElement(a)

}
