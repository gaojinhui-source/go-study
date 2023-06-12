package __collections

import (
	"fmt"
	"testing"
)

// 寻找最长不含有重复字符的子串
// abcabcbb -> abc
// bbbb -> b
// pwwkew -> kew

// 不能处理中文
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start, maxLength := 0, 0
	for i, v := range []byte(s) {
		// 如果这个字符在map中存在并且这个下标大于start，就把start重置为该字符的下标
		if val, ok := lastOccurred[v]; ok && val >= start {
			start = val + 1
		}

		// 如果子串的长度大于maxlength,则更细
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		// 正常情况下该map保存值和下标
		lastOccurred[v] = i
	}
	return maxLength
}

func TestMapExample(t *testing.T) {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
}
