package algo

import (
	"fmt"
	"sync"
)

//使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
//
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

// channel

func AlternatingPrintingofNumbersAndLetters() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true // 通知letter协程
			}
		}
	}()

	wait.Add(1) // 用于字母结束停止协程
	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true // 通知number协程
			}
		}
	}()

	// 启动 number 协程
	number <- true

	wait.Wait() // 阻塞协程，直到计数器归零
}

// 总结
// 两个协程，互相等待对方通知，letter 协程判断结束条件。
// 需要手动启动 number 协程
// 通过sync.WaitGroup{}来阻塞当前wait 所在goroutine，计数器归零后才能执行后续代码
