package algo

import (
	"fmt"
	"unicode"
)

// 有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，
//问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。
//可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。 问最后机器人的坐标是多少？ 注意 R123(LF)

const (
	Left = iota
	Top
	Right
	Bottom
)

func RobotCoordinate() {
	// 坐标初始化
	x0, y0 := 0, 0
	x1, y1, z1 := move("R2(LF)", x0, y0, Top) // Top 表示方向正Y
	fmt.Println(x1, y1, z1)
}

// x,y  z 代表当前方向（如正Y,负Y）
func move(cmd string, x0, y0, z0 int) (x1, y1, z1 int) {
	repeat := 0     // 记录重复次数
	repeatCmd := "" // 记录重复指令
	for _, v := range cmd {
		switch {
		case unicode.IsNumber(v):
			repeat = repeat*10 + int(v-'0') //** 计算位读数 比如123 = (1*10+2)*10+3
		case v == ')':
			// 递归调用重复指令
			for range repeat {
				x0, y0, z0 = move(repeatCmd, x0, y0, z0)
			}
			// 重置重复状态
			repeat = 0
			repeatCmd = ""
		case repeat > 0 && v != '(' && v != ')':
			repeatCmd = repeatCmd + string(v)
		case v == 'L':
			// 左转
			//** 改 z 调整方向 取模数
			z0 = (z0 - 1 + 4) % 4 // +4 避免负数取模
		case v == 'R':
			// 右转
			z0 = (z0 + 1) % 4
		case v == 'F':
			switch {
			case z0 == Left:
				x0 = x0 - 1
			case z0 == Right:
				x0 = x0 + 1
			case z0 == Top:
				y0 = y0 + 1
			case z0 == Bottom:
				y0 = y0 - 1
			}
		case v == 'B':
			switch {
			case z0 == Left:
				x0 = x0 + 1
			case z0 == Right:
				x0 = x0 - 1
			case z0 == Top:
				y0 = y0 - 1
			case z0 == Bottom:
				y0 = y0 + 1
			}
		}

	}
	return x0, y0, z0
}

// 总结： unicode.IsNumber() 重复指令，那么将重复次数和重复的指令存起来递归调用
