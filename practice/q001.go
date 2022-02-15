package main

import (
	"fmt"
	"sync"
)

// 交替打印数字和字母
//
//问题描述
//
//使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
//
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func main() {
	waitGroup := sync.WaitGroup{}
	letter, number := make(chan bool), make(chan bool)
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Println(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}

	}()
	waitGroup.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wait.Done()
					return
				}

				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
			}

		}
	}(&waitGroup)
	number <- true
	waitGroup.Wait()
}
