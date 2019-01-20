public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m = new Main();

        System.out.println(m.reverse(123));
    }
    public int reverse(int x){
        int result = 0;
        for(int i=0;x!=0;i++){
            int dig = x%10;
             int tmp = result*10 + dig;
            if((tmp-dig)/10 != result){
                return 0;
            }
            x/=10;
            result = tmp;
        }
        return result;
    }
}