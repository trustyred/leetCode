import java.lang.reflect.AnnotatedArrayType;
import java.lang.reflect.Array;
import java.text.CollationElementIterator;
import java.util.*;

public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m  = new Main();
        System.out.println(m.fourSum(new int[] {1,0,-1,0,-2,2},0));
    }
    public List<List<Integer>> fourSum(int [] nums, int target){
        HashMap map = new HashMap();
        HashSet set = new HashSet();
        for(int i=0;i<nums.length;i++){
            map.put(nums[i],i);
        }

        for(int i=0;i<nums.length;i++){
            for(int j=i+1;j<nums.length;j++){
                for(int k=j+1;k<nums.length;k++){
                    int dst = target - nums[i] - nums[j] -nums[k];
                    if(map.containsKey(dst)){
                        int find_key = (int) map.get(dst);
                        if(find_key != i && find_key != j && find_key!= k)
                        {
                            List<Integer> tmp = Arrays.asList(nums[i],nums[j],nums[k],nums[(int) map.get(dst)]);
                            Collections.sort(tmp);
                            set.add(tmp);
                        }

                    }
                }
            }
        }
        return new ArrayList(set);
    }
//    public List<List<Integer>> fourSum(int[] nums, int target) {
//        Arrays.sort(nums);
//
//        HashSet result = new HashSet();
//        for (int i = 0; i < nums.length; i++) {
//            for (int j = i + 1; j < nums.length; j++) {
//                int lo = 2;
//                int hi = nums.length - 1;
//                int dest = target - nums[i] - nums[j];
//                if(lo==i || lo == j || hi==i||hi==j){
//                    continue;
//                }
//                while (lo < hi) {
//
//                    int sum = nums[lo] + nums[hi];
//                    if (sum == dest) {
//                        List<Integer> l = Arrays.asList(nums[i], nums[j], nums[lo], nums[hi]);
//                        Collections.sort(l);
//                        result.add(l);
//                        lo++;
//                        hi--;
//                    } else if (sum < dest) {
//                        lo++;
//                    } else {
//                        hi--;
//                    }
//
//                }
//            }
//        }
//        return new ArrayList(result);
//    }

}
