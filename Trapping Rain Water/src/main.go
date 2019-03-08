package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	*list.List
}
func main() {
	//result := trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})
	result := trap([]int{4,2,3})

	fmt.Println(result)
}
func New() *Stack{
	return &Stack{
		list.New(),
	}
}
func (stack *Stack) push(elem interface{}){
	stack.PushBack(elem)
}
func (stack *Stack) pop() interface{}{
	if stack.Len() != 0{
		val := stack.Back()
		stack.Remove(val)
		return val.Value
	}else{
		return false
	}
}

func (stack *Stack) peak() interface{}{
	return stack.Back().Value
}

func (stack *Stack) isEmpty() bool{
	if stack.Len() == 0{
		return true
	}else{
		return false
	}
}
func (stack *Stack) clear(){
	stack.Init()
}
func trap(height []int) int {
	s := New()
	low := 0
	hi := 1
	if len(height) == 0{
		return 0
	}
	max := 0
	sum :=0
	for low<len(height)-1{
		if height[hi] >= height[low]{
			if s.isEmpty(){
				low = hi
				hi ++
			}else{
				basic := 0
				if height[hi] <= height[low]{
					basic = height[hi]
				}else{
					basic = height[low]
				}
				for ! s.isEmpty(){
					sum += basic - s.pop().(int)
				}
				low = hi
				hi ++
			}
		}else{

			s.push(height[hi])
			hi++
			if s.peak().(int) > max{
				max = s.peak().(int)
			}
		}
		if hi == len(height) && s.isEmpty(){
			if max == height[hi-1]{
				hi--
				basic := 0
				if height[hi] <= height[low]{
					basic = height[hi]
				}else{
					basic = height[low]
				}

				for ! s.isEmpty(){
					sum += basic - s.pop().(int)
				}
			}
			break
		}else if hi ==len(height) &&! s.isEmpty(){
			low++
			hi = low +1
			s.clear()
		}
	}
	return sum
}