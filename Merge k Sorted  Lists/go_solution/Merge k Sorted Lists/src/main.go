package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	//var ListNode_list []*ListNode
	head1 := create([] int {-1,1})
	head2 := create([] int {-7,-5,-3,1,4})
	head3 := create([] int {-2,-1,0,2})


	print_link(mergeKLists([]*ListNode{head1,head2,head3}))
}
func print_link(head *ListNode){
	p := head
	for p!=nil{
		fmt.Print(p.Val," ")
		p = p.Next
	}
}
func create(nums []int) *ListNode{
	var head,p,q *ListNode
	head,p,q = nil,nil,nil
	for i:=0;i<len(nums);i++  {
		if head == nil{
			head = new(ListNode)
			head.Next = nil
			head.Val = nums[i]
			p=head;
			q=head;
		} else {
			p = new(ListNode)
			p.Val = nums[i]
			q.Next = p
			p.Next = nil
			q=p
		}
	}
	return head
}


func LinkLastPointer(head *ListNode) *ListNode{
	p := head
	for p.Next!=nil{
		p = p.Next
	}
	return p

}
func mergeKLists(lists []*ListNode) *ListNode {
	interval := 1
	var head *ListNode
	var new_lists []*ListNode

	for i:=0;i<len(lists);i++{
		if lists[i] != nil{
			new_lists = append(new_lists, lists[i])
		}
	}
	if len(new_lists) == 1{
		return new_lists[0]
	}
	for interval<len(new_lists){
		for i:=0;i<len(new_lists)-interval;i+=2*interval{

			head = merge2List(new_lists[i], new_lists[i+interval])
			new_lists[i] = head
 		}
		interval *=2
	}
	return head

}

func merge2List(l1 *ListNode , l2 *ListNode) *ListNode{
	head := l1
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}

	p1,q1 := l1,l1
	p2 := l2

	for p1 != nil && p2 !=nil {
		if p1.Val > p2.Val {

			tmp := p2.Next
			if p1 == head{
				p2.Next = p1
				q1 = p2
				head = q1
			}else{
				p2.Next = p1
				q1.Next = p2
				q1 = p2
			}
			p2 = tmp
		}else{
			q1 = p1
			p1 = p1.Next
		}
	}
	if p2 != nil{
		q1.Next = p2
	}

	return head
}