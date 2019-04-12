package main

type ListNode struct{
	Val int
	Next *ListNode
}
func main() {

}


func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	q := head
	for p!=nil{
		if p != head{
			if q.Val == p.Val{
				q.Next = p.Next
				p = q
			}
		}
		q = p
		p=p.Next
	}
	return head
}
func deleteDuplicates2(head *ListNode) *ListNode {
	p := head
	q := head
	for p!=nil{
		if p != head{
			if q.Val == p.Val{
				q.Next = p.Next
				p = q
			}
		}
		q = p
		p=p.Next
	}
	return head
}
