package main

import (
	"fmt"
)

func main() {
	result := searchRange([]int{3,3,3,3},3)
	fmt.Println(result)
}
func searchRange(nums []int, target int) []int {
	result := []int{-1,-1}
	if len(nums) < 1{
		return result
	}
	leftIndex := twoDividedSearch(nums,target,true)

	if leftIndex == len(nums) || nums[leftIndex] != target{
		return result
	}
	rightIndex := twoDividedSearch(nums,target,false)

	result[0] = leftIndex
	result[1] = rightIndex - 1
	return result
}
func twoDividedSearch(nums []int,target int,left bool) int{
	lo := 0
	hi := len(nums) - 1
	var mid int

	if lo == hi && ! left{
		return lo+1
	}
	for lo<hi{
		mid = (lo+hi)/2
		if nums[mid] >target || (left && nums[mid] == target){
			hi = mid
		}else {
			lo = mid+1
		}
	}
	if nums[lo] == target && !left{
		lo ++
	}
	return lo
}