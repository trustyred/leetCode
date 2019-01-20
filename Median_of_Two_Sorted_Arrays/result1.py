#!/usr/bin/python
class Solution(object):
	def findMedianSortedArrays(self, nums1, nums2):
		nums1.extend(nums2)
		nums1.sort()
		length = len(nums1)
		if not length%2:
			return (nums1[length/2]+nums1[length/2-1])/2.0
		else:
			return nums1[length/2]/1.0
		



if __name__ == "__main__":
	obj = Solution()
	nums1 = [1,3]
	nums2 = [2]
	print obj.findMedianSortedArrays(nums1, nums2)

