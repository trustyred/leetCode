package main

import "fmt"

func main() {
	nums := []int{1}
	sortColors(nums)
	fmt.Println(nums)

}
func sortColors(nums []int){
	num0 := 0
	num1 := 0
	num2 := 0

	for i:=0;i<len(nums);i++{
		if nums[i] == 0{
			num0++
		}else if nums[i] == 1{
			num1++
		}else if nums[i] == 2{
			num2++
		}
	}
	for i:=0;i<len(nums);i++{
		if i<num0{
			nums[i] = 0
		}else if i<num0+num1{
			nums[i] = 1
		}else if i<num0+num1+num2{
			nums[i] = 2
		}
	}
}