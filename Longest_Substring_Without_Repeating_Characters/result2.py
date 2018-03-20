class Solution(object):
	def lengthOfLongestSubstring(self, s):
		i=0
		j=0
		ans=0
		length = len(s)
		hash_map = dict()
		while j<length:
			if hash_map.has_key(s[j]):
				i = max(i,hash_map[s[j]])
			ans = max(ans,j-i+1)
			hash_map[s[j]]=j+1
			j+=1
		return ans
            
if __name__ == "__main__":
	s = "abcaefghijklmnopq"
	obj = Solution()
	print obj.lengthOfLongestSubstring(s)
