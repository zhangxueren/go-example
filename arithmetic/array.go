package arithmetic

/**
两数之和 II - 输入有序数组
@link https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
@tags 双指针
 */
func twoSum(numbers []int, target int) []int {
	start,end := 0,len(numbers)-1
	for start <= end{
		if target - numbers[start] == numbers[end] {
			return []int{start+1, end+1}
		} else if target - numbers[start] > numbers[end] {
			start ++
		} else {
			end --
		}
	}
	return nil
}

/**
344. 反转字符串
https://leetcode-cn.com/problems/reverse-string/
@tags 双指针、原地赋值
 */
func ReverseString(s []byte)  {
	start, end := 0, len(s)-1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}