package main

import (
	"fmt"
)

func main() {
	result := combine(2,2)
	fmt.Println(result)
}

func combine(n int, k int) [][]int {
	var remain []int
	for i:=0;i<n;i++{
		remain = append(remain,i+1)
	}
	var result [][]int
	result = backtrack([]int{},remain,k,result)
	return result
}

func backtrack(current []int, remain []int,n int,result [][]int) [][]int{
	if n==0{
		//tmpCurrent := make([]int,len(current))
		//copy(tmpCurrent,current)
		result = append(result,current)
		return result
	}
	lengthRemain := len(remain)
	for i:=0;i<lengthRemain;i++{
		//fmt.Println("remain",remain)
		tmp := make([]int,len(remain))
		copy(tmp,remain)
		tmp = delRemain(tmp,0)
		tmpCurrent := make([]int,len(current))
		copy(tmpCurrent,current)
		//fmt.Println("tmp",tmp)

		result = backtrack(append(tmpCurrent,remain[0]),tmp,n-1,result)
		remain = delRemain(remain,0)
		//fmt.Println("remain2",remain)
	}
	return result
}
func delRemain(remain []int,index int) []int{
	remain = append(remain[:index],remain[index+1:]...)
	return remain
}