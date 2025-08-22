package algo

import (
	"fmt"
	"math/rand"
	"sync"
)

// 写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
//另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

// 满意答案需注意：
// goroutine 在golang中是非阻塞的
// channel 无缓冲情况下，读写都是阻塞的，且可以用for循环来读取数据，当管道关闭后，for 退出。
// golang 中有专用的select case 语法从管道读取数据。

func GoroutineAndChannel() {
	//byDone()
	byWaitGroup()
}

func byDone() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			num := rand.Intn(10)
			ch <- num
		}
		close(ch)
	}()

	go func() {
		for {
			select {
			case num, ok := <-ch:
				if ok {
					fmt.Println("num", num)
				} else {
					done <- true
				}
			}
		}
	}()

	<-done // 阻塞主程序
	close(done)
}

func byWaitGroup() {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(out)
		for i := 0; i < 5; i++ {
			num := rand.Intn(10)
			out <- num
		}
	}()

	go func() {
		defer wg.Done()
		for i := range out { // 新知识
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
