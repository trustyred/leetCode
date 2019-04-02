package main

import (
	"container/list"
	"fmt"
	"strings"
)

type Stack struct{
	list *list.List
}
func main() {
	result := simplifyPath("/a/../../b/../c//.//")
	fmt.Println(result)
	m := map[int]int{
		1:2,
	}
	e := m[2]
	fmt.Println(e)
}
func New() *Stack{
	return &Stack{list:list.New(),}
}
func (stack *Stack)Push(val interface{}){
	stack.list.PushBack(val)
}
func (stack *Stack)Pop() interface{}{
	if stack.list.Len() < 1{
		return nil
	}else{
		val := stack.list.Back()
		stack.list.Remove(val)
		return val.Value
	}
}
func (stack *Stack)Peek() interface{}{
	if stack.list.Len() < 1{
		return nil
	}else{
		val := stack.list.Back()
		return val.Value
	}
}
func (stack *Stack)show() []string{
	var s []string
	for e:=stack.list.Front();e!=nil;e=e.Next(){
		s = append(s,e.Value.(string))
	}
	return s
}

func simplifyPath_err(path string) string {
	s := New()
	if len(path) < 2{
		return path
	}

	pre := path[0]
	for i:=1;i<len(path);i++{
		tmp := s.Peek()
		//fmt.Printf("%c ",tmp)
		if path[i] == '/' && tmp != nil && tmp.(byte) != '/'{
			s.Push(path[i])
		}else if path[i] == '.' && pre == '.'{
			s.Pop()
			s.Pop()
		}else if path[i] !='.' && path[i] != '/'{
			s.Push(path[i])
		}
		//s.show()
		//fmt.Println()
		pre = path[i]

	}
	resultStack := New()
	result := s.Pop()
	for result != nil{
		resultStack.Push(result)
		result = s.Pop()
	}
	result = resultStack.Pop()
	var resultByteList []byte
	resultByteList = append(resultByteList,'/')
	for result != nil{
		resultByteList = append(resultByteList,result.(byte))
		result = resultStack.Pop()
	}

	if len(resultByteList) == 1{
		return string(resultByteList)
	}else{
		return string(resultByteList[:len(resultByteList)-1])
	}
	return string(resultByteList)
}
func simplifyPath(path string) string {
	skip := map[string]int{
		"":   1,
		".":  2,
		"..": 3,
	}
	r := strings.Split(path,"/")
	var processStrL []string
	for i:=range r{
		if len(r[i]) !=0 {
			processStrL = append(processStrL,r[i])
		}
	}
	s := New()
	for i:=0;i<len(processStrL);i++{
		if processStrL[i] == ".."{
			s.Pop()
		}else{
			val := skip[processStrL[i]]
			if val ==0{
				s.Push(processStrL[i])
			}
		}
	}
	result := s.show()
	a := strings.Join(result,"/")
	return "/" + a
}