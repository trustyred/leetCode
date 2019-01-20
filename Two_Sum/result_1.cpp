#include <iostream>
#include <vector>
using namespace std;
/*
	本题解为Two_Sum的暴力解法，也是最容易想到的

*/
class Solution {
	public:
		vector<int> twoSum(vector<int>& nums, int target){

                	vector<int> vec;
 			for(int i=0;i<nums.size();i++)
                	{
 				cout<<"start"<<nums.size()<<endl;
 				for(int j=i+1;j<nums.size();j++)
				{
					if(nums[i]+nums[j]==target)
					{
						cout<<"find"<<endl;
						vec.push_back(i);
						vec.push_back(j);
 						break;
					}
			
				}
			


			}
                        return vec;
			
		}


};
int main(void)
{
    vector<int> vec,vec2;
    int a[4] = {2,7,11,5};
    for(int i=0;i<4;i++){
        vec.push_back(a[i]);
    }
    Solution b;
    vec2 = b.twoSum(vec, 9);
    for(int i=0;i<vec2.size();i++){
        cout<<vec2[i]<<endl;
    }
    
    return 0;
}
