package main

import (
	"fmt"
)

func main() {
	result := getPermutation(4,1)
	fmt.Println(result)
	//result := factorial(600)
	//fmt.Println(result)
}

func getPermutation_old(n int, k int) string {
	var startByteList []byte
	if n==0{
		return ""
	}

	for i:=1;i<=n;i++{
		startByteList = append(startByteList, byte(i)+'0')
	}
	var result []string
	result = backtrack([]byte{}, string(startByteList),result,n)
	return result[k]
}

func backtrack(currentString []byte,  remain string, result []string,n int) []string{
	if len(remain) == 0{
		result = append(result,string(currentString))
		return result
	}
	for i:=0;i<len(remain);i++{
		tmp := make([]byte, len(currentString))
		tmp = append(tmp, remain[i])
		copy(tmp, currentString)
		remainString := delStringByte(remain, i)
		result = backtrack(tmp, remainString,result,n)
	}
	return result

}
func delStringByte(s string, index int) string{
	byteList := []byte(s)
	tmp := make([]byte, len(s))
	copy(tmp, byteList)
	tmp = append(tmp[:index],tmp[index+1:]...)
	return string(tmp)
}

func getPermutation(n int, k int) string {
	var serial []byte
	for i:=1;i<=n;i++{
		serial = append(serial,byte('0'+i))
	}
	var result []byte
	for len(result)<n{
		facResult := factorial(n-1)
		index := k/facResult
		//将第一个元素添加到结果中
		result = append(result, serial[index])
		//删除掉这个元素
		serial = append(serial[:index],serial[index+1:]...)
		k -= index*facResult
	}
	return string(result)
}

func factorial(n int) int{
	if n == 0{
		return 1
	}
	return n*factorial(n-1)
}