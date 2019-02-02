import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

class Main {
    HashMap<String,String> map = new HashMap<String,String>(){
        {
            put("2","abc");
            put("3","def");
            put("4","ghi");
            put("5","jkl");
            put("6","mno");
            put("7","pqrs");
            put("8","tuv");
            put("9","wxyz");
        }
    };
    List<String> list = new ArrayList<>();


    private void backtrack(String comm, String next){

        if(next.length() == 0){
            list.add(comm);
            return ;
        }
        String cur_dig = next.substring(0,1);
        String cur_dig_str = map.get(cur_dig);

        for(int i=0;i<cur_dig_str.length();i++){

            backtrack(comm + cur_dig_str.charAt(i),next.substring(1));
        }


    }

    public List<String> letterCombinations(String digits) {
        if(digits.length()==0){
            return list;
        }
        backtrack("",digits);
        return list;

    }

    public static void main(String[] args) {
        //每一步只关注怎么把当前的字符串和下一个字符添加到一起
        Main m = new Main();
        System.out.println(m.letterCombinations("23"));
    }
}