package algo

import (
	"fmt"
	"strings"
)

// 给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
// 这里规定【大小写为不同字符】，且考虑字符串重点空格。
// 给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。 保证两串的长度都小于等于5000。

func IsEqualsBeforeSortString() {
	fmt.Println(isEqualsBeforeSortString("123 2 32", "123 23 2"))
}

// 首先要保证字符串长度小于5000。
// *之后只需要一次循环遍历s1中的字符在s2是否都存在即可。
// 遍历s1字符时顺带检查s2中同样的字符
func isEqualsBeforeSortString(s1 string, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}
	
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
