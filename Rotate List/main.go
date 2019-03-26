package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func main() {

	head := create([]int{})
	head.print()
	fmt.Println()
	result := rotateRight(head,10)

	result.print()
}

func create(l []int) *ListNode{
	var head *ListNode
	var p,q *ListNode

	for i:=0 ;i<len(l);i++{
		if head == nil{
			p = new(ListNode)
			p.Val = l[i]
			p.Next = nil
			head = p
			q = p
		}else{
			p = new(ListNode)
			p.Val = l[i]
			p.Next = nil
			q.Next = p
			q = p
		}
	}
	return head
}
func (l *ListNode) print(){
	p := l
	for p!=nil{
		fmt.Printf("%d,",p.Val)
		p = p.Next
	}
}
func rotateRight(head *ListNode, k int) *ListNode {

	var p,q, newTail, newHead *ListNode
	p=head
	q=p
	i:=0
	listLen := 0
	for p!=nil{
		listLen++
		p=p.Next
	}
	if listLen == 0{
		return head
	}
	k %=listLen
	p=head
	for  p!=nil{
		if i == listLen-k{
			newTail = q
			newHead = p
		}
		q=p
		p=p.Next
		i++
	}
	if k==0{
		return head
	}
	q.Next = head
	newTail.Next = nil
	return newHead
}