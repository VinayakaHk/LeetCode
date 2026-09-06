package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		_, ok := set[v]
		if ok {
			print("4. ", true)
			return true
		} else {
			set[v] = true
		}
	}
	print("4. ", false)
	return false
}
