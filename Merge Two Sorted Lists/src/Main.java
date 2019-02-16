
public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m = new Main();
    }

        public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
            ListNode head=l1;
            ListNode p1 = l1,q1=l1,p2=l2,q2=l2;
            if(l1==null){
                return l2;
            }
            if(l2==null){
                return l1;
            }
            while(p1!=null && p2!=null){
                ListNode tmp;
                if(p1.val > p2.val){
                    tmp = p2.next;
                    if(p1==head){
                        p2.next = p1;
                        head = p2;
                        q1 = head;
                    }
                    else{
                        p2.next = p1;
                        q1.next = p2;
                        q1 = p2;
                    }
                    p2=tmp;

                }
                else if(p1.val <= p2.val){
                    q1=p1;
                    p1=p1.next;
                }
            }
            if(p2!=null){
                q1.next = p2;
            }
            return head;
        }

}
