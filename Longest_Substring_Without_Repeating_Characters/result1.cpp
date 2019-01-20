#include <iostream>
#include <string>
#include <set>
using namespace std;
/*
本题目使用滑动窗口的方法来解决对每个子串的遍历问题
从而使复杂度变为了o(n)
初始化hashset为空
每次判断当前窗口的最右边是否在hashset中
如果不在 说明当前的窗口为一个不重复子串，进行长度计算，并且将当前窗口最右边加入hashset
如果在 说明当前窗口子串出现重复了，那么就将窗口左边进行移动，并且将窗口左边的值在hashset中删除

其实从这里看出这个算法还有优化的余地，因为每次删除的只是窗口最左边的值，而并不一定是重复的值
   */
class Solution{
	public:
		int lengthOfLongestSubstring(string s) {
			set<int> hash_set;
			int i,j,max_num=0;
			for(i=0,j=0;i<s.length()&&j<s.length();){
				if(hash_set.find(s[j])==hash_set.end()){
					hash_set.insert(s[j++]);
					max_num = max_num>(j-i)?max_num:(j-i);
				}
				else{
					hash_set.erase(s[i++]);
				}
				
			}
			return max_num;
		}
	

};
int main(void){
	int num;
	string s="bbbbbbbb";
	Solution obj = Solution();
	num = obj.lengthOfLongestSubstring(s);
	cout<<num<<endl;

	return 0;

}
