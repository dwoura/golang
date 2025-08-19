package algo

import (
	"fmt"
	"strings"
	"unicode"
)

// 请编写一个方法，将字符串中的空格全部替换为“%20”。 假定该字符串有足够的空间存放新增的字符，
//并且知道字符串的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。
//给定一个string为原始的串，返回替换后的string。

func StringReplace() {
	stringReplace("123a")
	stringReplace("abc CED cdg")
}

func stringReplace(s string) {
	r := []rune(s)
	if len(r) > 1000 {
		fmt.Println("字符串不符合规则，小于1000")
		return
	}
	for _, v := range r {
		if string(v) != " " && unicode.IsLetter(v) == false {
			fmt.Println("字符串不符合规则")
			return
		}
	}

	s = strings.Replace(s, " ", "%20", -1)
	fmt.Println(s)
}

// 总结： unicode.IsLetter() strings.Replace() 两个方法
