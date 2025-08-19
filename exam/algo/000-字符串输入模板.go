package algo

import (
	"bufio"
	"fmt"
	"os"
)

func stringInput() {
	// 字符串输入模板
	scan := bufio.NewScanner(os.Stdin)
	s := ""
	for scan.Scan() {
		s = scan.Text()
	}
	fmt.Println(s)
}
