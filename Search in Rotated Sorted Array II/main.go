package main

import (
	"fmt"
)

func main() {
	result := search([]int{1,1,3},1)
	fmt.Println(result)
}
func search(nums[]int, target int)bool{
	lo := 0
	hi := len(nums)-1
	for lo<=hi{
		for lo<hi && nums[lo] == nums[hi]{
			lo++
		}
		mid := (lo+hi)/2
		if target == nums[mid]{
			return true
		}
		if nums[mid] > nums[lo]{
			if target < nums[mid] && nums[lo] <= target{
				hi = mid
			}else{
				lo++
			}
		}else if nums[mid] < nums[lo]{
			if target > nums[mid]  && nums[hi] >= target{
				lo++
			}else{
				hi = mid
			}
		}else{
			lo++
		}
	}
	return false
}
