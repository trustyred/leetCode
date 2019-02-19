import org.w3c.dom.ls.LSException;

public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m = new Main();
        ListNode head = m .create(new int[] {});
//        m.print_link(head);
        head = m.swapPairs(head);
        m.print_link(head);
    }
    public void print_link(ListNode head){
        ListNode p = head;
        while (p!=null){
            System.out.println(p.val);
            p = p.next;
        }

    }

    public ListNode create(int[] nums) {
        ListNode head = null;
        ListNode p, q = null;
        for (int i = 0; i < nums.length; i++) {
            if (head == null) {
                p = new ListNode(nums[i]);
                p.next = null;
                head = p;
                q = head;
            } else {
                p = new ListNode(nums[i]);
                q.next = p;
                q = p;
            }
        }
        return head;
    }

    public ListNode swapPairs(ListNode head) {
        ListNode p,q,f=null;
        p = head;
        if(head == null){
            return head;
        }
        while(p.next!=null){
            if(p==head){
                q = p.next;
                p.next = q.next;
                q.next = p;
                head = q;
                f=p;
            }
            else{
                q = p.next;
                p.next = q.next;
                q.next = p;
                f.next = q;
                f=p;
            }

            p=p.next;
            if(p==null){
                return head;
            }
        }
        return head;
    }
}
