package main

import (
	"fmt"
	"strings"
)

func main() {
	result := lengthOfLastWord("h ")
	fmt.Println(result)
}
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	sl:=strings.Split(s," ")
	for i:=len(sl)-1;i>=0;i--{
		if len(sl[i]) != 0{
			return len(sl[i])
		}
	}
	return len(sl[len(sl)-1])
}