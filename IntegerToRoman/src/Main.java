import java.util.HashMap;
import java.util.Map;

public class Main {
    public static void main(String[] args) {
        Main m = new Main();
        int [] area_data = {1,8,6,2,5,4,8,3,7};
        System.out.println(m.intToRoman(1994));
    }
    public String intToRoman(int num) {
        int [] roman_list = {1000,900,500,400,100,90,50,40,10,9,5,4,1};
        StringBuilder romanNum = new StringBuilder();
        Map numToRoman = new HashMap();
        numToRoman.put(1000,"M");
        numToRoman.put(900,"CM");
        numToRoman.put(500,"D");
        numToRoman.put(400,"CD");
        numToRoman.put(100,"C");
        numToRoman.put(90,"XC");
        numToRoman.put(50,"L");
        numToRoman.put(40,"XL");
        numToRoman.put(10,"X");
        numToRoman.put(9,"IX");
        numToRoman.put(5,"V");
        numToRoman.put(4,"IV");
        numToRoman.put(1,"I");

        int tmp=0;
        int tmp1=0;
        while(num !=0 ){
            for(int i=0;i<roman_list.length;i++){
                if(num >= roman_list[i]){
                    tmp = roman_list[i];
                    break;
                }
            }
            romanNum.append(numToRoman.get(tmp));
            num -= tmp;
        }

        return romanNum.toString();
    }



}