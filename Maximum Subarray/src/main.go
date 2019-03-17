package main

import "fmt"

func main() {
	result := maxSubArray([]int{1,2})
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	dp := make([]int,len(nums))
	max := nums[0]
	dp[0] = nums[0]
	for i:=1;i<len(nums);i++{
		addNum := 0
		if dp[i-1] >0{
			addNum = dp[i-1]
		}else{
			addNum = 0
		}
		dp[i] = nums[i] + addNum
		if dp[i] > max{
			max = dp[i]
		}
	}
	return max
}


