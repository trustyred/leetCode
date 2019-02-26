package main

import "fmt"

func main() {
	//fmt.Println(longestValidParentheses(")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"))
	fmt.Println(longestValidParentheses("((()))"))

}

func longestValidParentheses(s string) int {
	var dp [99999999]int
	if len(s) == 0 {
		return 0
	}
	dp[0] = 0
	for i:=1;i<len(s);i++{
		if i == 1{
			if s[0] == '(' && s[1] == ')'{
				dp[1] = 2
			}else{
				dp[1] = 0
			}
		}else if s[i] == ')' && s[i-1] == '('{
			dp[i] = dp[i-2] +2
		}else if s[i] == ')' && s[i-1]==')' && s[i-dp[i-1]-1]=='('{
			dp[i] = dp[i-1] +2 + dp[i-dp[i-1]-2]
		}else if s[i]=='('{
			dp[i] = dp[i-1]
		}
	}
	return dp[len(s)-1]
}