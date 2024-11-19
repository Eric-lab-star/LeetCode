package main

import "fmt"

func main() {
	an := searchInsert([]int{1,3,5,6}, 5)
	fmt.Println("an", an)
}

func searchInsert(nums []int, target int) int {

	res := 0;
	for _, v := range nums {
		if v >= target {
			return res;
		}
		res++
	}
	return res;
}
