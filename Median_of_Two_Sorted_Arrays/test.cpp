#include <iostream>
#include <vector>
using namespace std;
int max(int a,int b);

int min(int a, int b);

int max(int a, int b){
	return a>b?a:b;
}
int min(int a, int b){
	return a>b?b:a;
}
class Solution{
	public:
		double findMedianSortedArrays(vector<int> &nums1, vector<int> &nums2){
			vector<int> temp;
			if(nums1.size()<nums2.size()){
				temp = nums1;
				nums1 = nums2;
				nums2 = temp;
			}
			int left_max=0,right_min=0;
			int m = nums1.size();
			int n = nums2.size();
			int halflen = (m+n+1)/2;
			int j;
			for(int i=0;i<m;i++){
				j = halflen-i;
				if((nums1[i-1]<=nums2[j]) && (nums1[i]>=nums2[j-1])){
					if(i==0){
						left_max = nums2[j-1];
					}	
					else if(j==0){
						left_max = nums1[i-1];
					}
					else{
						left_max = max(nums1[i-1],nums2[j-1]);
					}

					if(i==m){
						right_min = nums2[j];
					}
					else if(j==m){
						right_min = nums1[i];
						
					}
					else{
						right_min = min(nums1[i],nums2[j]);
					}
					
				
				}
			
			
			
			}
			return (left_max+right_min)/2.0;
		}


};
int main(void){
	int a1[2] = {1,3};
	int a2[2] = {2};
	vector<int> nums1,nums2;
	Solution obj = Solution();
	for(int i=0;i<1;i++){
		nums1.push_back(a1[i]);
	}
	for(int i=0;i<2;i++){
		nums2.push_back(a2[i]);
	}
	cout<<obj.findMedianSortedArrays(nums1, nums2)<<endl;
	
	return 0;
}
