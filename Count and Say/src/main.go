package main

import (
	"strings"
	"fmt"
)

func main() {
	var a []int
	a = make([]int,10)
	a[4] = 5
	fmt.Println(countAndSay(7))

}
func countAndSay(n int) string {
	//var basicMap map[int]string = map[int]string{1:"11",11:"21",2:"12"}
	sequence := make([]string,25)
	sequence = append([]string{"1","11","21","1211","111221"},sequence...)
	if n<6{
		return sequence[n-1]
	}
	for i:= 5;i<n;i++{
		var tmp []string
		lastSequence := sequence[i-1]
		var last byte
		constantVal := lastSequence[0]
		constantTime := 1
		for j:=0;j<len(lastSequence);j++{
			if j==0{
				last = lastSequence[j]
			}else{
				if last == lastSequence[j]{
					constantVal = lastSequence[j]
					constantTime ++

				}else{
					tmp = append(tmp,fmt.Sprintf("%d%d",constantTime,constantVal-'0'))
					constantTime = 1
					constantVal = lastSequence[j]
				}
				last = lastSequence[j]
				if j == len(lastSequence)-1{
					tmp = append(tmp,fmt.Sprintf("%d%d",constantTime,constantVal-'0'))
				}
			}
		}
		sequence[i] = strings.Join(tmp,"")
	}
	return sequence[n-1]
}
