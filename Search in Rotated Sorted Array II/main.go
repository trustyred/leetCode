package main

import (
	"fmt"
)

func main() {
	result := search([]int{1,3},3)
	fmt.Println(result)
}
func search(nums[]int, target int)bool{
	var realMid int
	if len(nums) ==0{
		return false
	}
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := (lo+hi)/2
		if (nums[mid] < nums[0]) && (target < nums[0])||((nums[mid] > nums[0]) && (target >nums[0])){
			realMid = nums[mid]
		}else if target < nums[0]{
			realMid = -653524738
		}else {
			realMid = 653524738
		}

		if target > realMid {
			lo ++
		}else if target < realMid {
			hi = mid
		}else{
			return true
		}
	}
	if nums[(lo+hi)/2] == target{
		return true
	}
	return false
}
