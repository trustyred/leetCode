package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int {1,2,3,4,5,2},2))
}
func removeElement(nums []int, val int) int {
	n := len(nums)
	i := 0
	for i<n {
		if nums[i] == val{
			nums[i] = nums[n-1]
			n--
		}else{
			i++
		}
	}

	return n
}
