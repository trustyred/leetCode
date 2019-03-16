package main

func main() {

}

func maxSubArray(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	max := nums[0]
	maxIndex := 0
	for i:=1;i<len(nums);i++{
		if nums[i] > max{
			maxIndex = i
			max = nums[i]
		}
	}

}