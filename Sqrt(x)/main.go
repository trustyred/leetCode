package main

import "fmt"

func main() {
	result := mySqrt(100)
	fmt.Println(result)
}
func mySqrt(x int) int {
	result := x
	for result*result>x{
		result = (result+x/result)/2
	}
	return int(result)
}