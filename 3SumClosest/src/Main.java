import java.util.Arrays;
import java.util.HashSet;

public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
        Main m = new Main();
        System.out.println(m.threeSumClosest(new int[] {0,1,2},3));
    }
            public int threeSumClosest(int[] nums, int target) {
        //整体的思路和3Sum一致，主要就是加了一个最小值的记录功能
                int tmp = 0;
                int result = 0 ;
                boolean isfind = false;
                int min=Integer.MAX_VALUE;
                Arrays.sort(nums);
                for(int i=0;i<nums.length;i++){
                    int lo = i+1;
                    int hi = nums.length-1;
                    int dest = nums[i];
                    while(lo<hi){
                        tmp = nums[lo] + nums[hi] + dest;

                        int sum = target - (nums[lo] + nums[hi]);
                        if(sum == dest){
                            isfind = true;
                            result = tmp;
                            break;
                        }
                        else if(sum>dest){
                            if(Math.abs(tmp - target) < min){
                                min = Math.abs(tmp -target);
                                result = tmp;
                            }
                            lo++;
                        }
                        else{
                            if(Math.abs(tmp - target) < min){
                                min = Math.abs(tmp -target);
                                result = tmp;
                            }
                            hi--;
                        }
                    }
                    if(isfind){
                        break;
                    }
                }
                return result;
            }
}
