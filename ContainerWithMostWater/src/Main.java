import java.util.Arrays;
import java.util.Collections;

public class Main {
    public static void main(String[] args) {
        Main m = new Main();
        int [] area_data = {1,8,6,2,5,4,8,3,7};
        int result = m.maxArea(area_data);
        System.out.println(result);

    }

    public int maxArea(int[] height) {
        int max = 0;
        for(int i=0,j=height.length-1 ;i!=j;){
            int minLine = (int) Collections.min(Arrays.asList(height[i],height[j]));
            int currentContainer = minLine * (j-i);
            if(currentContainer > max){
                max = currentContainer;
            }

            if(height[i] < height[j]){
                i++;
            }
            else{
                j--;
            }
        }
        return max;

    }

}