#include <iostream>
using namespace std;

struct ListNode {
	int val;
	ListNode *next;
	ListNode(int x) :val(x), next(NULL) {}
};
class Solution{

	public:


};

int main(void){
	struct ListNode a(10);
	a = ListNode(5);
	
	cout<<a.val<<endl;

	return 0;
}
