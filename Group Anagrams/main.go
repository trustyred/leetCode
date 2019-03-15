package main

import (
"sort"
"fmt"
)

type byteSlice []byte

func (byteslice byteSlice) Len() int{
	return len(byteslice)
}

func (byteslice byteSlice) Less(i,j int) bool{
	return byteslice[i] < byteslice[j]
}
func (byteslice byteSlice) Swap(i,j int){
	byteslice[i],byteslice[j] = byteslice[j],byteslice[i]
}

func groupAnagrams(strs []string) [][]string {
	result := map[string] []string{}


	for i:=0;i<len(strs);i++{
		bl := byteSlice(strs[i])
		sort.Sort(bl)
		str := string(bl)
		if val,err := result[str];err == false{
			var tmp []string
			tmp = append(tmp, strs[i])
			result[str] = tmp
		}else{
			var tmp[]string
			tmp = append(val,strs[i])
			result[str] = tmp
		}
	}
	var r [][]string
	for _,val := range result{
		r = append(r,val)
	}
	return r
}



func main() {
	r := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(r)
}
