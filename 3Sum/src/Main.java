import java.net.Inet4Address;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m = new Main();
        System.out.println(m.threeSum(new int[] {-1, 0, 1, 2, -1, -4}));
    }
    public List<List<Integer>> threeSum(int[] nums) {
        //关键就在这个排序，让查找的时间复杂度降了下来
        Arrays.sort(nums);
        List<List<Integer>> result = new ArrayList<List<Integer>>();
        for(int i=0;i<nums.length;i++){
            int hi = nums.length - 1;
            int lo = i+1;
            int target = - nums[i];
            while(lo<hi){
                int sum = nums[lo] + nums[hi];
                if(sum == target){
                    result.add(Arrays.asList(nums[i],nums[lo],nums[hi]));
                    while(lo<hi && nums[lo] == nums[lo+1]){
                        lo++;
                    }
                    lo++;
                    while(hi>lo && nums[hi] == nums[hi-1]){
                        hi--;
                    }
                    hi--;
                }
                else if(sum > target){
                    lo++;
                }
                else{
                    hi--;
                }
            }

        }
        return result;
    }
}
