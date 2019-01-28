public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        String a = "abcd";
        System.out.println(a.substring(0));
        Main m = new Main();

        System.out.println(m.isMatch("cccaab","cccaab*"));
    }
    public boolean isMatch(String text, String pattern) {
        boolean[][] dp = new boolean[text.length() + 1][pattern.length() + 1];
        dp[text.length()][pattern.length()] = true;

        for (int i = text.length(); i >= 0; i--){
            for (int j = pattern.length() - 1; j >= 0; j--){
                if(i<text.length() && j+1<pattern.length()){
                    dp[i][j] = (dp[i+1][j+1] &&(text.charAt(i) == text.charAt(j) || pattern.charAt(i) == '.'))||(pattern.charAt(j+1) == '*'&&(dp[i][j+2]|| (((pattern.charAt(j) == text.charAt(i) || pattern.charAt(j) == '.'))) && dp[i+1][j]));
                }
                else{
                    dp[i][j] = ((i < text.length() && (pattern.charAt(j) == text.charAt(i) || pattern.charAt(j) == '.'))) && dp[i+1][j+1];
                }

            }
        }
        return dp[0][0];
    }
}
