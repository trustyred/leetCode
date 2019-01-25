import javax.swing.table.TableRowSorter;

public class Main {
    public static void main(String[] args) {
        System.out.println("hello");
        Main m = new Main();
        System.out.println(m.isPalindrome(1221));

    }
    public boolean isPalindrome(int x) {
        if(x<0){
            return false;
        }
        int save_x = x;
        int reverse_x = 0;
        while(x!=0){
            int num = x%10;
            reverse_x = 10*reverse_x + num;
            x/=10;
        }
        if(reverse_x == save_x){
            return true;
        }
        else{
            return false;
        }

    }
}
