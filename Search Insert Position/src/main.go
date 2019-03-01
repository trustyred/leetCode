package main

import (
	"fmt"
)

func main() {
	result := searchInsert([]int{1,2,3},4)
	fmt.Println(result)
}
func searchInsert(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	var mid int
	if len(nums) == 0{
		return 0
	}
	for lo < hi {
		mid = (lo+hi)/2
		if nums[mid] > target{
			hi = mid
		}else if nums[mid] < target{
			lo++
		}else{
			return mid
		}
	}
	mid = (lo+hi)/2
	if target > nums[mid]{
		return mid + 1
	}else{
		return mid
	}
}
