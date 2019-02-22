package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divideBit(1,1))
}
//solution not use bit calc
func divide(dividend int, divisor int) int {
	var result = 0
	var flag = 1
	if (divisor<0 && dividend>0) ||divisor>0 && dividend<0 {
		flag = -1
	}
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	for dividend >= divisor{
		dividend -= divisor
		result ++
	}
	result = result * flag
	if result> 2147483647{
		return 	2147483647
	}else if result < -2147483648{
		return -2147483648
	}
	return result
}
//soluttion use bit calc
// Runtime:4ms
// Memory Usage:2.4 MB
func divideBit(dividend int, divisor int) int {
	var result = 0
	var flag = 1
	if (divisor<0 && dividend>0) ||divisor>0 && dividend<0 {
		flag = -1
	}
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	sub := divisor
	c := 1
	for dividend >= divisor{
		if dividend >= sub{
			dividend -= sub
			result += c
			sub = sub <<1
			c = c<<1
		}else{
			sub = sub >>1
			c = c>>1
		}
	}
	result = result * flag
	if result> 2147483647{
		return 	2147483647
	}else if result < -2147483648{
		return -2147483648
	}
	return result
}