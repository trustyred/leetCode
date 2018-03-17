#-*- coding: UTF-8 -*-
#!/usr/bin/python 

#Two_Sum的python解法，经过调整后的时间复杂度为O(n)
#理解本思路，要抓住题干中只有一个确定解，也就是输出一定只有两个下标数据
#以题目中数组nums = [2,11,15,7] target = 9为例子
#第一次循环9 - 2 = 7此时hash_table还没有值，跳过
#第二次循环9 -11 = -2 不符合if条件，跳过
#第三次循环 9 - 15 = -6 不符合if条件，跳过
#第四次循环 9 - 7 =2 符合条件，返回函数
#我们使用这种后添加hash_table的方法，一定会错过一次真正要输出的下标
#但是错过之后，下次碰到另一个符合加和为target的数时，就会立刻判断
#成功。

class Solution(object):
	def twoSum(self, nums, target):
		hash_table = dict()
		for k,v in enumerate(nums):
			flag = target - v
			if hash_table.has_key(flag):
				return [k,hash_table[flag]]
			hash_table[v] = k




if __name__ == "__main__":
	l = [2,7,11,15]
	obj = Solution()
	print obj.twoSum(l, 9)
