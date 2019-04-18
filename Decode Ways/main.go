package main

import (
	"fmt"
)

func main() {
	result := numDecodings("0")

	fmt.Println(result)
}
func numDecodings(s string) int {
	if len(s) == 0{
		return 0
	}else{
		if s[0] == '0'{
			return 0
		}
	}

	dp := make([]int,len(s))
	for i:=0;i<len(s);i++{
		if i==0{
			if s[i] == '0'{
				dp[i] = 0
			}else{
				dp[i] = 1
			}
		}else if i>0 && s[i] != '0' && (s[i-1]>='0'&&s[i-1]<'3') && (s[i]>'0'&&s[i]>'7'){
			/*
			1 : 1
			1,2: 1,2 12
			1,2,2: 1,2,2 12,2 1,22
			1,2,2,2: 1,2,2,2 12,2,2 1,22,2 1,2,22 12,22
			1,2,2,2,2: 1,2,2,2,2 12,2,2,2 1,22,2,2 1,2,22,2 12,22,2 1,2,2,22 12,2,22 1,22,22
			 */
			dp[i] =

		}
	}

	fmt.Println(dp)
	return dp[len(s)]
}
