package main

import "fmt"

func main() {
	nums1 := []int{1,2,4,5,6,0}
	nums2 := []int{3}
	merge(nums1,5,nums2,1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i := 0
	j := 0
	c := 0
	largeNum := 0
	if m == 0{
		for k:=0;k<n;k++{
			nums1[k] = nums2[k]
		}
		return
	}else if n==0{
		return
	}
	for i<m && j<n{
		if nums1[c] >= nums2[j]{
			for k := m+c-largeNum;k>c;k--{
				nums1[k] = nums1[k-1]
			}
			nums1[c] = nums2[j]
			c++
			j++
		}else{
			i++
			c++
			largeNum ++
		}
	}
	if j<n{
		for k := m+j;k<m+n;k++{
			nums1[k] = nums2[j]
			j++
		}
	}
}
