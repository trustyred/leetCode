public class Main {
    public static void main(String[] args) {
        Main m = new Main();
        int result = m.myAtoi("2147483648");
//        System.out.println(-2147483648-10);
//        System.out.println(-2147483648-8);
        System.out.println(result);

    }
    public int myAtoi(String str){
        int result = 0;
        str = str.trim();
        int flag = 1;
        int c_num=0,tmp_result=0;
        for(int i=0;i<str.length();i++){
            char c= str.charAt(i);
            //根据第一位的符号来修改计算逻辑的乘数，为正就乘1，为负就乘-1
            if(i==0){
                if(c=='+'){
                    continue;
                }
                else if(c=='-'){
                    flag = -1;
                    continue;
                }
            }
//走到这里发现任何非数字字符的符号，全部进行返回操作
            if(c>'9' || c<'0'){
                return result;
            }

            c_num = c-'0';
            tmp_result = result;
            result = 10*result + c_num*flag;
//          判断过界的逻辑，未过界的数除以10得到之前的结果，但是一旦过界后这个数再除以10一定会与原本的数符号不一致，导致最终不相等
            if(result/10!=tmp_result){

                return flag==1? 2147483647:-2147483648;
            }

        }

        return result;

    }
}