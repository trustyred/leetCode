import java.util.HashMap;
import java.util.HashSet;
import java.util.Stack;

public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m = new Main();
        boolean result = m.isValid("((()))[]");
        System.out.println(result);
    }
    public boolean isValid(String s) {
        Stack stack = new Stack<Character>();
        HashSet left_brackets = new HashSet();
        HashSet right_brackets = new HashSet();
        HashMap match_brackets = new HashMap();
        left_brackets.add('(');
        left_brackets.add('[');
        left_brackets.add('{');
        right_brackets.add(')');
        right_brackets.add(']');
        right_brackets.add('}');
        match_brackets.put('(',')');
        match_brackets.put('[',']');
        match_brackets.put('{','}');

        char tmp;
        for (char c:s.toCharArray()
             ) {
            if(left_brackets.contains(c)){
                stack.push(c);
            }
            else if (right_brackets.contains(c)){
                if(stack.isEmpty()){
                    return false;
                }
                tmp = (char) stack.pop();
                char match_right = (char) match_brackets.get(tmp);
                if(match_right != c){
                    return false;
                }
            }
        }
        if(stack.isEmpty()){
            return true;
        }
        else{
            return false;
        }

    }
}
