package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{1,3},3))
}
func search(nums[]int, target int)int{
	var real_mid int
	var result int
	if len(nums) ==0{
		return -1
	}
	result = -1
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := (lo+hi)/2
		if (nums[mid] < nums[0]) == (target < nums[0]){
			real_mid = nums[mid]
		}else if target < nums[0]{
			real_mid = -653524738
		}else {
			real_mid = 653524738
		}

		if target > real_mid{
			lo ++
		}else if target < real_mid{
			hi = mid
		}else{
			result = (lo+hi)/2
			return result
		}
	}
	if nums[(lo+hi)/2] == target{
		return (lo+hi)/2
	}
	return result
}

func search_binary_error(nums []int, target int) int{
	var middle int
	middle = -1
	l := 0
	r := len(nums) - 1
	result := -1
	if len(nums) == 0{
		return result
	}
	for l!=r&&(l>=0&&l<len(nums)-1)&&(r>=0&&r<=len(nums)-1){
		if (l+r)/2 == middle{
			l++
			middle = (l+r)/2
		}else{
			middle = (l+r)/2
		}

		if nums[middle] < target {
			if middle == len(nums) -1 {
				break
			}
			if nums[middle+1] >= nums[middle] && nums[middle]>nums[middle-1]{
				l = middle
			}else {
				r = middle
			}
		}else if nums[middle] > target {
			if nums[l] <= target {
				r = middle
			}else {
				l = middle
			}
		}else{
			result = middle
			break
		}
	}
	if l == r && nums[(l+r)/2] == target {
		return l
	}
	return result
}
func backtrack(){

}