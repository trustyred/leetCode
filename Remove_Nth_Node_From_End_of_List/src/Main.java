import java.util.List;

public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m = new Main();
        ListNode head;
        int [] arr = new int[] {1};
        head = m.createLinkList(arr);
        head = m.removeNthFromEnd(head,1);
        ListNode p = head;
        while(p!=null){
            System.out.println(p.val);
            p = p.next;
        }

    }
    public ListNode createLinkList(int [] arr){
        ListNode p=null,q=null,head=null;
        for(int i=0;i<arr.length;i++){
            if(i==0){
                head = new ListNode(arr[i]);
                p=head;
                q=head;
            }
            else{
                q=p;
                p= new ListNode(arr[i]);
                q.next = p;
            }
            p.next=null;
        }
//        p=head;
//        for(int i=0;i<arr.length;i++){
//            p=p.next;
//        }
        return head;
    }

    public ListNode removeNthFromEnd(ListNode head, int n) {
            ListNode p=head,q=head;
            int num = 0 ;

            while(p!=null){
                p=p.next;
                num++;
            }

            p=head;
            for(int i=0;i<num-n;i++){
                q=p;
                p=p.next;
            }

            if(num == n ){
                head = head.next;
            }
            else{
                q.next = p.next;
            }

            return head;
    }
}
