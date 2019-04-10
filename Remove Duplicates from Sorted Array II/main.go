package main

import (
	"fmt"
)

func main() {
	a := []int{0,0,1,1,1,1,2,3,3}
	result := removeDuplicates(a)
	fmt.Println(result)
	fmt.Println(a)
}

func removeDuplicates(nums []int) int {
	if len(nums)<3{
		return len(nums)
	}
	currentNum := nums[0]
	delNum := 0
	duplicateNum :=1
	numLength := len(nums)
	for i:=1;i<numLength;i++{
		if nums[i] == currentNum{
			duplicateNum ++

		}else{
			if duplicateNum>2{
				delNum += duplicateNum - 2
				i = forwardMove(nums,i,duplicateNum-2)
				numLength -= duplicateNum - 2
			}
			currentNum = nums[i]
			duplicateNum = 1
		}
	}
	if duplicateNum >2{
		fmt.Println(duplicateNum)
		delNum += duplicateNum - 2
		forwardMove(nums,numLength,duplicateNum-2)
		numLength -= duplicateNum - 2
	}
	//fmt.Println(delNum)
	//fmt.Println(nums)
	return len(nums) - delNum
}
func forwardMove(nums []int, start int,bits int) int{
	for i:=start;i<len(nums);i++{
		nums[i-bits] = nums[i]
	}

	nums = nums[:len(nums)-bits]
	return start - bits
}