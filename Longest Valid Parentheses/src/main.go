package main

import (
	"container/list"
	"fmt"
)
type Stack struct {
	list *list.List
}
func main() {
	//fmt.Println(longestValidParentheses(")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"))
	//fmt.Println(longestValidParentheses(")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"))
	fmt.Println(longestValidParentheses("((()))"))
}
func NewStack() *Stack{
	l := list.New()
	stack := Stack{l}
	return &stack
}
func (stack *Stack)Push(value interface{}){
	stack.list.PushBack(value)
}
func (stack *Stack)Pop() interface{}{
	e := stack.list.Back()
	if e!= nil{
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}
func (stack *Stack)Empty() bool{
	return stack.list.Len() == 0
}
func (stack *Stack)Peek() interface{}{
	e := stack.list.Back()
	if e != nil{
		return e.Value
	}
	return nil
}
func longestValidParentheses(s string) int{
	stack := NewStack()
	max := 0
	stack.Push(-1)
	for i:=0;i<len(s);i++{
		if s[i] == '('{
			stack.Push(i)
		}else if s[i] == ')'{
			stack.Pop()
			if stack.Empty(){
				stack.Push(i)
			}else{
				if i-stack.Peek().(int) > max{
					max = i - stack.Peek().(int)
				}
			}
		}

	}

	return max
}
func longestValidParentheses_dp(s string) int {
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