package main

import (
	"fmt"
)

func main() {
	result := multiply("60974249908865105026646412538664653190280198809433017", "500238825698990292381312765074025160144624723742")
	//result := multiply("2", "3")

	fmt.Println(result)
}

func multiply(num1 string, num2 string) string {
	n := len(num1)
	m := len(num2)
	if (m == 1 && num2[0] == '0') || (n == 1 && num1[0] == '0') {
		return "0"
	}
	sum := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := (num1[j] - '0') * (num2[i] - '0')
			low := mul % 10
			hi := mul / 10
			sum[i+j] += int(hi)
			sum[i+j+1] += int(low)
		}
	}

	sumBytes := make([]byte, m+n)
	for i := len(sum) - 1; i > 0; i-- {
		if sum[i] > 9 {
			sum[i-1] += sum[i] / 10
			sum[i] = sum[i] % 10
		}
		sum[i] += '0'
		sumBytes[i] = byte(sum[i])
	}
	sum[0] += '0'
	sumBytes[0] = byte(sum[0])
	var s string
	for i := 0; i < len(sum); i++ {
		if sum[i] != '0' {
			//fmt.Println(i)
			s = string(sumBytes[i:])
			break
		}
	}

	return s
}
