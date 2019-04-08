package main

import (
	"fmt"
)

func main() {
	result := subsets([]int{})
	fmt.Println(result)
}
func subsets(nums []int) [][]int {
	var result [][]int
	result = backtrack([]int{},nums,result)
	result = append(result,[]int{})
	return result
}
func backtrack(current []int, remain []int,result [][]int) [][]int{

	if len(remain) == 0 {
		return result
	}
	remainLength := len(remain)
	for i:=0;i<remainLength;i++{
		tmpCurrent := make([]int,len(current))
		copy(tmpCurrent,current)
		tmpCurrent = append(tmpCurrent,remain[0])
		result = append(result,tmpCurrent)
		tmpRemain := make([]int,len(remain))
		copy(tmpRemain,remain)
		tmpRemain = append(tmpRemain[:0],tmpRemain[1:]...)
		fmt.Println(tmpRemain)
		result = backtrack(tmpCurrent, tmpRemain,result)
		remain = append(remain[:0],remain[1:]...)

	}
	return result
}

