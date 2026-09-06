package main

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
	print("3. ", maxV)
	return maxV
}
