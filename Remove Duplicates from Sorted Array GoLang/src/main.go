package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([] int{0,0,0,1,1,1,2,2,3,3,4}))
}
func removeDuplicates(nums []int) int {
	j := 0
	for i:=1;i<len(nums);i++{
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j+1
}
