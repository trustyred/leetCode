package main

import (
	"sort"
	"fmt"
)

func main() {
	a := []int{2,3,1}
	nextPermutation(a)
	fmt.Println(a)
}
func nextPermutation(nums []int)  {
	lastIndex := len(nums) - 1
	selectIndex := 0
	isMax := true
	for i:=len(nums)-2;i>=0;i--{
		cur := nums[i]
		if cur < nums[lastIndex]{
			isMax = false
			min := 653524738
			for j:= lastIndex;j<len(nums);j++{
				if nums[j] > nums[i] && nums[j]-nums[i]<min{
					selectIndex = j
				}
			}
			tmp := nums[selectIndex]
			nums[selectIndex] = nums[i]
			nums[i] = tmp
			sort.Ints(nums[i+1:])
			break
		}
		lastIndex = i
	}
	if isMax{
		sort.Ints(nums)
	}
}
