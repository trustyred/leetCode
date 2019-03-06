package main

import (
	"fmt"
	"sort"
)


func main() {

	result := combinationSum2([]int{10,1,2,7,6,1,5},8)
	fmt.Println(result)
}

func combinationSum2(candidates []int, target int) [][]int {
	var set [][]int
	sort.Ints(candidates)
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
			if i>start && candidates[i] == candidates[i-1]{
				continue
			}
			tmp = append(tmp,candidates[i])
			set = backtrack(tmp, candidates,set,target,remain-candidates[i],i+1)
			//fmt.Println(tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	return set

}