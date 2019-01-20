public class Main {

    public static void main(String[] args) {
        Main a = new Main();

        String maxStr = a.longestPalinndrome("aaaa");
        System.out.println(maxStr);
    }
    private String longestPalinndrome(String s){
        String strMax="";
        boolean [] [] dpArr = new boolean[1000][1000];
        for(int i=s.length();i>=0;i--){
            for(int j=i ;j<s.length();j++){
                dpArr[i][j] = (s.charAt(i) == s.charAt(j)) && ((j-i<3) || dpArr[i+1][j-1]);

                if(dpArr[i][j] &&(strMax.equals("") ||j-i+1>strMax.length())){
                    strMax = s.substring(i,j+1);
                }
            }
        }

        return strMax;
    }
    public Boolean isPalinndrome(String s){
        String reverseString = new StringBuilder(s).reverse().toString();
        return reverseString.equals(s);
    }
}
