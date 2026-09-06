package main

import "fmt"

func print(num_of_program string, a any) {
	fmt.Println(num_of_program, a)
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	target := 9
	twoSum(a, target)
	bestTimeToBuyAndSellStock(a)
	majorityElement(a)
	containsDuplicate(a)
	moveZeroes(a)
	sortedSquares(a)
}
