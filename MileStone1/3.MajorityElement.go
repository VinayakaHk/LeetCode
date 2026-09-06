package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	maxMap := make(map[int]int)
	maxV := 0
	for _, v := range nums {
		_, ok := maxMap[v]
		if ok {
			maxMap[v] = maxMap[v] + 1
		} else {
			maxMap[v] = 1
		}
	}
	for i, v := range maxMap {
		if v > len(nums)/2 {
			maxV = i
		}
	}
	fmt.Println(maxV)
	return maxV
}
