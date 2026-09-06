package main

import "fmt"

func print(a any) {
	fmt.Println(a)
}

func main() {
	a := []int{0, 0, 1}
	// target := 9
	// twoSum(a, target)
	// bestTimeToBuyAndSellStock(a)
	// majorityElement(a)
	// containsDuplicate(a)

	moveZeroes(a)
}
