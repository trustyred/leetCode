package main

import (
	"fmt"
	"math"
)

func main() {
	result := addBinary("1111","11")
	fmt.Println(result)
}
func addBinary(a string, b string) string {
	aByteList := []byte(a)
	bByteList := []byte(b)
	zeroList := make([]byte,int(math.Abs(float64(len(b)-len(a)))))
	for i:=0;i<len(zeroList);i++{
		zeroList[i] = byte('0')
	}
	if len(a) < len(b){
		aByteList = append(zeroList,aByteList...)
	}else{
		bByteList = append(zeroList,bByteList...)
	}
	carry := byte(0)
	for i:=len(aByteList)-1;i>=0;i--{
		aByteList[i] = aByteList[i]-'0' + bByteList[i] + carry
		carry = (aByteList[i]-'0')/2
		aByteList[i] = ((aByteList[i] - '0') % 2) + '0'
	}
	if carry != 0{
		tmp := []byte{carry+'0'}
		aByteList = append(tmp,aByteList...)
	}
	return string(aByteList)
}
