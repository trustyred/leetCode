package main

import (
	"strings"
	"fmt"
)

func main() {
	var result []int
	result = findSubstring("foobarfoobar",[]string{"foo","bar"})
    fmt.Println(result)
}

func RemoveIndex(s []string, index int) []string {
	nowStr := make([]string,len(s))
	copy(nowStr,s)
	return append(nowStr[:index], nowStr[index+1:]...)
}

//solution use backtrack but Time Limit Exceeded
func findSubstring(s string, words []string) []int {
	substrToIndex := map[int]string{}
	if len(s) == 0 || len(words)==0{
		return []int{}
	}
	var result []int
	backtrack(s,substrToIndex,"",words)
	if len(substrToIndex) == 0{
		return []int{}
	}else{
		fmt.Println(substrToIndex)
		for i := range substrToIndex{
			result = append(result, i)
		}
		return result
	}

	return []int{1,2,3}
}
func backtrack(strBase string,substrToIndex map[int]string,cur string , words []string){
	if len(words) ==0 {
		var index int
		index = 1
		i := 0
		for index != -1 {
			index := strings.Index(strBase[i:], cur)
			//fmt.Println(strBase[i:],cur,index,i)

			if index != -1{
				substrToIndex[index+i] = cur
				i++
			}else{
				break
			}
		}
		return
	}
	for i:= range words{
		//fmt.Println(words,words[i])
		curWords := RemoveIndex(words,i)
		backtrack(strBase, substrToIndex, cur + words[i], curWords)
	}

}
