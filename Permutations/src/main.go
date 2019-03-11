package main

import (
	"fmt"
)

func main() {
	result := permute([]int{})
	fmt.Println(result)
}

func permute(nums []int) [][]int {
	result := backtrack([]int{},nums,[][]int{},len(nums))
	return result
}

func backtrack(current []int,remain []int, result [][]int, n int)[][]int{
	if len(current) == n{

		result = append(result, current)
		return result
	}
	for i:=0;i<len(remain);i++{
		c := append(current,remain[i])
		r := sliceDel(remain,i)
		result = backtrack(c,r,result,n)
	}
	return result
}

func sliceDel(l []int,num int) []int{
	tmp := make([]int,len(l))
	copy(tmp, l)
	tmp = append(tmp[:num],tmp[num+1:]...)
	return tmp
}