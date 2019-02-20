import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m = new Main();
        ListNode head = m.create(new int[]{1});
        head = m.reverseKGroup(head,3);
        m.print_link(head);
    }
    public ListNode create(int[] nums){
        ListNode head,p,q;
        head = null;
        q = null;
        p = null;
        for(int i=0;i<nums.length;i++){
            if(head == null){
                p = new ListNode(nums[i]);
                p.next = null;
                head = p;
                q = p;
            }
            else{
                p = new ListNode(nums[i]);
                p.next = null;
                q.next = p;
                q = p;
            }
        }
        return head;
    }
    public void print_link(ListNode head){
        ListNode p = head;
        while (p!=null){
            System.out.println(p.val);
            p = p.next;
        }
    }
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode p = head;
        ListNode q = null;
        ListNode new_head = null;
        int len=0;
        len = ListLen(head);
        if(len==0 || len ==1){
            return head;
        }
        if(k>len){
            return head;
        }
        for(int i=0;i<len;i++){

            if(len-i < k){
                break;
            }
            if(i%k == 0){
                if(i==0){
                    new_head = reverseList(p, k);
                    p = new_head;
                }
                else{
                    p = reverseList(p,k);
                    q.next = p;
                }
            }
            q = p;
            p=p.next;
        }
        return new_head;

    }
    public int ListLen(ListNode head){
        ListNode p = head;
        int len = 0;
        while (p!=null){
            len++;
            p=p.next;
        }
        return len;
    }
    public ListNode reverseList(ListNode head, int k){
        ListNode p = head;
        ListNode q = null;
        ListNode f = null;
        ListNode tail = null;
        for(int i=0;i<k;i++){
            if(p == head){
                q = p.next;
                p.next = null;
                tail = p;
                f = p;
                p=q;

            }
            else{
                q = p.next;
                p.next = f;
                f = p;
                p=q;
            }
        }

        head = f;
        tail.next = q;
        return head;
    }
    public int linkLength(ListNode head){
        ListNode p=head;
        int len=0;
        while (p!=null){
            p=p.next;
            len++;
        }
        return len;
    }
}
