package main

import (
	"fmt"
)

func main() {
	result := combinationSum([]int{2,3,5},7)
	//fmt.Println("---")
	fmt.Println(result)
}


func combinationSum(candidates []int, target int) [][]int {
	var set [][]int
	set = backtrack([]int{}, candidates,set,target,target,0)
	return set
}
func backtrack(tmp []int,candidates []int,set [][]int,target int,remain int,start int) [][]int{
	if remain <0{
		return set
	}else if remain == 0 {
		//fmt.Println(tmp)
		newL := make([]int,len(tmp))
		copy(newL, tmp)
		set = append(set, newL)
		return set
	}else{
		for i:=start;i<len(candidates);i++{
			tmp = append(tmp,candidates[i])
			set = backtrack(tmp, candidates,set,target,remain-candidates[i],i)
			//fmt.Println(tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	return set

}



//var set [][]int
//func combinationSum(candidates []int, target int) [][]int {
//
//	for i:=0;i<len(candidates);i++{
//		backtrack([]int{},candidates[i:],target)
//	}
//	return set
//}
//
//func sum(arr []int) int{
//	sum:= 0
//	for i:= range arr{
//		sum += arr[i]
//	}
//	return sum
//
//}
//
////问题所在就是current与set公用同一个底层数组了,我承认我是个被Python惯坏了的孩子
//func backtrack(current []int,remain []int,target int){
//	//fmt.Println(current)
//	s := sum(current)
//	if s == target{
//		newL := make([]int,len(current))
//		copy(newL, current)
//		set = append(set , newL)
//		return
//	}else if s > target{
//		return
//	}else{
//		for i:=0;i<len(remain);i++ {
//			current = append(current, remain[i])
//			//fmt.Println(current)
//			backtrack(current, remain[i:], target)
//		}
//	}
//
//}