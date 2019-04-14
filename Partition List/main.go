package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1,3,-1,5,2,1}
	fmt.Println(nums)
	head := create(nums)
	head = partition(head,3)
	printLink(head)
}

/*
1->4->3->2->5->2, x = 3
1->2->2->4->3->5
*/
func printLink(head *ListNode){
	p := head
	for p!=nil{
		fmt.Printf("%d ",p.Val)
		p = p.Next
	}
}
func create(nums []int) *ListNode{
	var head,p,q *ListNode

	for i:=0;i<len(nums);i++{
		p=new(ListNode)
		p.Next = nil
		p.Val = nums[i]
		if i==0{
			head = p
			q=head
		}else{
			q.Next = p
			q=p
		}
	}
	return head
}

/*
1->4->3->2->5->2, x = 3
1->2->2->4->3->5

3->4->1->1->3 x=3
1->1->3->4->3
*/
func partition(head *ListNode, x int) *ListNode {
	p := head
	q := head
	var t,tmp *ListNode
	for p!=nil{

		if p.Val < x && p!=head{
			if t == nil{
				tmp = p.Next
				p.Next = head
				head = p
				p = tmp
				q.Next = tmp
				t = head
			}else if t!=nil && t.Next !=p{
				tmp = p.Next
				p.Next = t.Next
				t.Next = p
				t=p
				p = tmp
				q.Next = tmp
			}else if t!=nil && t.Next ==p{
				t = p
				q = p
				p = p.Next
			}
		}else if p.Val<x && p==head{
			t=p
			q = p
			p = p.Next
		}else{
			q=p
			p=p.Next
		}
	}
	return head
}
