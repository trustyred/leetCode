import org.omg.CORBA.INTERNAL;

public class Main {

    public static void main(String[] args) {
        Main m = new Main();
        String s = m.longestCommonPrefix(new String[] {"c","acc","ccc"});
        System.out.println(s.length());
    }

    public String longestCommonPrefix(String[] strs) {
        int min= Integer.MAX_VALUE;
        if(strs.length==0){
            return "";
        }
        for(int i=0;i<strs.length;i++){
            if(strs[i].length() < min){
                min = strs[i].length();
            }
        }
        int i=0,j=0;
        char last_char='0';
        for(i=0;i<min;i++){
            boolean flag=true;
            for(j=0;j<strs.length;j++){
                if(j==0){
                    last_char = strs[j].charAt(i);
                }
                else{

                    if(last_char != strs[j].charAt(i)){
                        flag = false;
                        break;
                    }
                    last_char = strs[j].charAt(i);
                }

            }
            if(!flag){
                break;
            }
        }
        return strs[0].substring(0,i);
    }

}
