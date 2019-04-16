package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(grayCode(3))

}

func binToGray(bin int) int{
	return (bin>>1) ^ bin
}
func grayCode(n int) []int {
	i := 0
	var result []int
	for i<int(math.Pow(float64(2),float64(n))){
		result = append(result, binToGray(i))
		i++
	}
	return result
}
