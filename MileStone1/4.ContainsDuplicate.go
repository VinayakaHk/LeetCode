package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		_, ok := set[v]
		if ok {
			fmt.Println(true)
			return true
		} else {
			set[v] = true
		}
	}
	fmt.Println(false)

	return false
}
