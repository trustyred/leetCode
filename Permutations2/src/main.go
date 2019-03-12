package main

import (
	"fmt"
	"sort"
)

func main() {
	result := permuteUnique([]int{-1,2,-1,2,1,-1,2,1})
	fmt.Println(result)
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := backtrack([]int{},nums,[][]int{},len(nums))
	return result
}

func backtrack(current []int,remain []int, result [][]int, n int)[][]int{
	if len(current) == n{
		//这里真的是golang里面的坑，真是不长记性，按照这种添加方法，我的result是会变化的，因为result与current共享了
		//底层数组，这种问题还终极不好排查
		tmp := make([]int,len(current))
		copy(tmp,current)
		result = append(result, tmp)
		return result
	}
	for i:=0;i<len(remain);i++{
		if i<len(remain)-1 && remain[i+1] == remain[i]{
			continue
		}
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