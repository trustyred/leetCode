package main

import (
	"fmt"
)

func main() {
	result := combine(4,2)
	fmt.Println(result)
}

func combine(n int, k int) [][]int {
	var remain []int
	for i:=0;i<n;i++{
		remain = append(remain,i+1)
	}
	var result [][]int
	result = backtrack([]int{},remain,n,result)
	return result
}

func backtrack(current []int, remain []int,n int,result [][]int) [][]int{
	if n==0{
		tmpCurrent := make([]int,len(current))
		copy(tmpCurrent,current)
		result = append(result,tmpCurrent)
		return result
	}
	for i:=0;i<len(remain);i++{
		tmp := make([]int,len(remain))
		copy(tmp,remain)
		delRemain(tmp,i)
		result = backtrack(append(current,remain[i]),tmp,n-1,result)
	}
	return result
}
func delRemain(remain []int,index int){
	remain = append(remain[:index],remain[index+1:]...)
}