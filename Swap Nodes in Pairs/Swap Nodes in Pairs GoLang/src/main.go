package main

func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode{
	var p,q,f *ListNode
	p,q,f = head,head,head
	if head ==nil{
		return head
	}

	for p.Next != nil{
		if p == head{
			q = p.Next
			p.Next = q.Next
			q.Next = p
			head = q
			f = p
		}else{
			q = p.Next
			p.Next = q.Next
			q.Next = p
			f.Next = q
			f = p
		}
		p = p.Next
		if p == nil{
			return head
		}
	}
	return head
}
