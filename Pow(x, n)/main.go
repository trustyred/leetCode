package main

import (
	"fmt"
	"math"
)

func main() {
	result := myPow2(2,3)
	fmt.Println(result)
}
func myPow(x float64, n int) float64 {
	result := float64(1.0)
	if math.Abs(float64(x)) == float64(1.0000){
		if x<0 && n%2==0{
			result = math.Abs(x)
		}else if x<0{
			result = x
		}
		return result
	}
	for i:=0;i<int(math.Abs(float64(n)));i++{
		result *= x
	}
	if x<0 && n%2==0{
		result = math.Abs(result)
	}
	if n<0{
		result = 1.0/result
	}
	return result
}
/*速度最快的一个方法，优化了暴力方法的时间复杂度从O(n)到了log(n),每次都除以2这样就像是二分查找一样*/
func myPow2(x float64, n int) float64{
	if n == 0{
		return 1
	}
	if n<0{
		n = -n
		x = 1/x
	}
	if n%2 ==0 {
		return myPow2(x*x, n/2)
	}else{
		return x*myPow2(x*x,n/2)
	}
}