import java.util.ArrayList;
import java.util.List;

public class Main {

    public static void main(String[] args) {

        Main s = new Main();
        String ret = s.convert("ab",1);
        System.out.println(ret);
    }

    public String convert(String s ,int numRows){
        List<StringBuilder> rows = new ArrayList<>();
        if(numRows ==1){
            return  s;
        }

        for(int i=0;i<numRows;i++){
            rows.add(new StringBuilder());
        }
//        false为向下 true为向上
        boolean direction = true;
        int line = 0;

        for(char c:s.toCharArray()){
            if(line == 0 || line == numRows -1){
                direction = !direction;
            }
            rows.get(line).append(c);
            line += direction ? -1:1;
        }
        StringBuilder ret = new StringBuilder();
        for(StringBuilder row:rows){
            ret.append(row);
        }

        return ret.toString();

    }

}
