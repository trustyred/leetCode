import java.util.HashMap;
import java.util.Map;

public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m = new Main();
        System.out.println(m.romanToInt("MCMXCIV"));

    }
    public int romanToInt(String s) {
        int tmp=0;
        int sum=0;
        Map map =  new HashMap();
        map.put('I',1);
        map.put('V',5);
        map.put('X',10);
        map.put('L',50);
        map.put('C',100);
        map.put('D',500);
        map.put('M',1000);

        for(int i=s.length()-1;i>=0;i--){
            int current_num  = (int) map.get(s.charAt(i));
            if(current_num >= tmp){
                sum += current_num;
            }
            else{
                sum -= current_num;
            }
            tmp = current_num;
        }
        return sum;
    }
}
