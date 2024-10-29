package main

import "fmt"

func main() {
	nums := []int{
		1,
		3,
		2,
		5,
		2,
	}

	k := removeDuplicates(nums);
	fmt.Println(k);
}

func removeDuplicates(nums []int) int {
	prev := 0

	for _, num := range nums {
		if nums[prev] != num {
			prev++
			nums[prev] = num
		}
	}
	return prev + 1
}
