package main

import "fmt"

func main() {
	result := plusOne([]int{1,0,0})
	fmt.Println(result)
}
func plusOne(digits []int) []int {
	var carry int
	carry = 1
	for i:=len(digits)-1;i>=0;i--{
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
	}

	var result []int
	if carry != 0 {
		result = make([]int,len(digits)+1)
		copy(result[1:],digits)
		result[0] = carry
	}else{
		result = make([]int,len(digits))
		copy(result,digits)
	}
	return result
}