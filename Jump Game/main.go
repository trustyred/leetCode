package main

import (
	"fmt"
)

func main() {

	result := canJump([]int{0,2,3})
	fmt.Print(result)
}
func canJump(nums []int) bool {
	dp := make([]bool,len(nums))
	dp[0] = true
	for i:=1;i<len(nums);i++{
		for j:=i;j<len(nums)&& dp[i-1]&&j<i+nums[i-1];j++{
			dp[j] = true
		}
		if dp[len(nums)-1]{
			break
		}
	}
	return dp[len(nums)-1]

}
func canJump1(nums []int) bool {
	if len(nums) == 0 {
		return false
	}else if len(nums) == 1{
		return true
	}
	result := backtrack(0,nums,[]int{},[][]int{})
	state := false

	for _,i:=range result{
		if i[len(i)-1] != 0{
			state = true
			break
		}
	}
	return state
}

func backtrack(curr int, nums []int, currentPath []int,result [][]int) [][]int{
	if nums[curr] >= len(nums)-1-curr ||len(nums)-1 == curr || nums[curr] == 0 {
		currentPath = append(currentPath,nums[curr])
		tmp := make([]int,len(currentPath))
		copy(tmp, currentPath)
		result = append(result, tmp)
		return result
	}
	maxStep := nums[curr]
	currentPath = append(currentPath, nums[curr])
	for s:=1;s<=maxStep;s++{
		result = backtrack(curr+s,nums, currentPath,result)
	}
	return result
}