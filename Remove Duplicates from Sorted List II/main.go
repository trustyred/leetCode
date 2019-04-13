package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func main() {
	nums := []int{}
	head := create(nums)
	printLink(head)
	head = deleteDuplicates(head)
	fmt.Println()
	printLink(head)
}
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
func deleteDuplicates2(head *ListNode) *ListNode {
	p,q,t := head,head,head
	var delFlag *ListNode
	delFlag = nil
	delHead := false
	for p != nil{
		if p!=head{
			if p.Val == q.Val && q==t{
				q.Next = p.Next
				delHead = true
				q = t
				p = q
			}else if p.Val == q.Val{
				q.Next = p.Next
				delFlag = t
				q = t
				p = q
			}else{
				if delHead{
					head = head.Next
					q = head
					t = head
					delHead = false
				}else if delFlag != nil{
					delFlag.Next = p.Next
					q = t
					p = q
					delFlag = nil
				}
			}
		}
		t=q
		q=p
		p=p.Next
	}
	return head
}
func deleteDuplicates(head *ListNode) *ListNode{
	p := head
	q := head
	var t *ListNode
	t = nil
	duplicated := false
	for p!=nil{
		if p!=head{
			if p.Val != q.Val{
				if t == nil && !duplicated{
					t = q

				}else if t==nil && duplicated{
					head = p
					q = p
					duplicated = false
				}else if t!=nil && duplicated{
					t.Next = p
					duplicated = false
				}else {
					t = q
				}
			}else{
				duplicated = true
			}
		}
		q = p
		p = p.Next
	}
	if duplicated && t!=nil{
		t.Next = p
	}else if duplicated && t==nil{
		head = nil
	}
	return head
}