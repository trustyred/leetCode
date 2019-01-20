#include <iostream>
using namespace std;

/*
本题为使用链表实现大整数加法，根据本题了解到了c++中结构体可以声明，定义函数
并且对于函数可以使用冒号的方式进行初始化



   */
struct ListNode {
	int val;
	ListNode *next;
	ListNode(int x) :val(x), next(NULL) {}
};
class Solution{

	public:
		ListNode *addTwoNumbers(ListNode *l1, ListNode *l2){
			ListNode *p1,*p2,*head=NULL,*p=NULL,*q;
			int carry = 0;
			int sum = 0;

			for(p1=l1,p2=l2;p1!=NULL||p2!=NULL;){
				int p1_val,p2_val;
				if(p1){
					p1_val = p1->val;
				}
				else{
					p1_val=0;
				}

				if(p2){
					p2_val = p2->val;
				}
				else{
					p2_val=0;
				}

				sum = p1_val+p2_val+carry;
				carry = sum/10?1:0;
				sum %=10;
				if(!head){
					
					head = new ListNode(sum);
					q = head;
				}
				else{
					
					p = new ListNode(sum);
					q->next = p;
					q=p;
				}			
				if(p1){
				
					p1=p1->next;
				
				}
				if(p2){
				
					p2=p2->next;
				
				}
			
			}
			if(carry){
				p = new ListNode(carry);
				q->next = p;
			}
			return head;
		}

};
ListNode *create(int data[],int len){
	ListNode *p=NULL,*q=NULL,*head=NULL;
	for(int i=0;i<len;i++){
		if(head == NULL){
			p = head = new ListNode(data[i]);
			q = head;
	
		}
		else{
			p = new ListNode(data[i]);
			q->next = p;
			q=p;
		}
	}
	return head;
}
void print(ListNode *head){
	ListNode *p=head;
	while(p!=NULL){
		cout<<p->val<<endl;
		p=p->next;
	}
}

int main(void){
	int a1[100] = {5};
	int a2[100] = {5};
	ListNode *head1,*head2;
	head1 = create(a1,1);
	head2 = create(a2,1);
	
	Solution so = Solution();
	head1=so.addTwoNumbers(head1, head2);
	print(head1);
	

	return 0;
}
