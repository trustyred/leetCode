package main

import (
	"sort"
	"fmt"
)

func main() {

	result := firstMissingPositive([]int{1})
	fmt.Println(result)
}

func firstMissingPositive(nums []int) int {

	sort.Ints(nums)
	var target int
	if len(nums) == 0{
		return 1
	}

	for i:=1;i<len(nums);i++{
		if nums[i-1]>0 && nums[i] - nums[i-1] > 1{
			target = nums[i-1]+1
			return target
		}else if (nums[i-1] <=0 && nums[i] > 1) || nums[0] >1{
			target = 1
			return target
		}
	}
	if len(nums) >1 && nums[len(nums)-1] > 0{
		target = nums[len(nums) -1 ]+1
	}else{
		if len(nums) == 1 && nums[0] ==1{
			return 2
		}else{
			return 1
		}
	}
	return target
}
